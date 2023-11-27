package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/marcopeocchi/valeera/internal/config"
	"github.com/marcopeocchi/valeera/internal/models"
	"github.com/marcopeocchi/valeera/internal/utils"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

func Top(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	cpu, _ := cpu.Info()
	host, _ := host.Info()
	vmstat, _ := mem.VirtualMemory()

	err := json.NewEncoder(w).Encode(models.SystemTop{
		CPU:       cpu[0].ModelName,
		CoreCount: cpu[0].Cores + 1,
		Hostname:  host.Hostname,
		OS:        host.OS,
		Platform:  host.Platform,
		Uptime:    host.Uptime,
		RAMFree:   vmstat.Available,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Temp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	temp, _ := utils.ReadCPUTemp()

	err := json.NewEncoder(w).Encode(models.SystemTemp{
		CPUTemp: temp,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Config(c *config.SafeConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c.JsonEncoder(w)
	}
}

func Login(c *config.SafeConfig, rdb *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		user := models.User{}
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if user.Name != c.Cfg.Username && user.Password != c.Cfg.Password {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.Authenticated = true

		sessionID := uuid.NewString()
		ttl := time.Minute * 30

		err := rdb.Set(r.Context(), sessionID, user, ttl).Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "valeera_session",
			Value:    sessionID,
			Path:     "/",
			HttpOnly: true,
			Expires:  time.Now().Add(ttl),
		})
	}
}
