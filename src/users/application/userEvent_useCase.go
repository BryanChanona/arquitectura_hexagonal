package application

import (
	"sync"
)



type UserEvent struct {
	mu sync.Mutex
	listeners []chan struct{}
}

var userEvent = &UserEvent{}
// NewNotifyUser crea una nueva instancia de EventUser y notifica a todos los listeners.
func (userEvent *UserEvent) NewNotifyUser() {
	userEvent.mu.Lock()
	defer userEvent.mu.Unlock()

	// Notifica a todos los listeners
	for _, ch := range userEvent.listeners {
		ch <- struct{}{} // Envía una señal a cada listener
	}
	// Limpia la lista de listeners después de notificar
	userEvent.listeners = nil

	// Devuelve una nueva instancia de EventUser
}

// Wait permite a un caller registrarse para ser notificado cuando ocurra un evento.
func (e *UserEvent) Wait() <-chan struct{} {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Crea un canal con buffer para evitar bloqueos
	ch := make(chan struct{}, 1)

	// Agrega el canal a la lista de listeners
	e.listeners = append(e.listeners, ch)

	// Devuelve el canal para que el caller pueda esperar la notificación
	return ch
}

// NotifyUserUpdate notifica a todos los listeners que ha ocurrido una actualización de usuario.
func NotifyUserUpdate() {
	userEvent.NewNotifyUser()
	
}

func WaitForUserUpdate() <-chan struct{}{
	return userEvent.Wait()
}