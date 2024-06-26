package server

/*
func MiddlewareRateLimiting(next http.Handler, requestsPerSecond int) http.Handler {
	limiter := rate.NewLimiter(rate.Limit(requestsPerSecond), requestsPerSecond)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				log.Printf("Recovered fron panic %v", err)
			}
		}()
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			log.Println("Too many requests from", r.RemoteAddr)
			return
		}
		log.Println("Request allowed from", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
*/
