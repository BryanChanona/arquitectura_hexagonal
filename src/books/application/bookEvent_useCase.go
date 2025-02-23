package application

import (
	"sync"
)

type BookEvent struct {
	mu sync.Mutex
	listeners []chan struct{}
}

var bookEvent = &BookEvent{}
// NewNotifyUser crea una nueva instancia de EventUser y notifica a todos los listeners.
func (bookEvent *BookEvent) NewNotifyBook() {
	bookEvent.mu.Lock()
	defer bookEvent.mu.Unlock()

	// Notifica a todos los listeners
	for _, ch := range bookEvent.listeners {
		ch <- struct{}{} // Envía una señal a cada listener
	}
	// Limpia la lista de listeners después de notificar
	bookEvent.listeners = nil

	// Devuelve una nueva instancia de EventUser
}

// Wait permite a un caller registrarse para ser notificado cuando ocurra un evento.
func (e *BookEvent) Wait() <-chan struct{} {
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
func NotifyBookUpdate() {
	bookEvent.NewNotifyBook()
	
}

func WaitForBookUpdate() <-chan struct{}{
	return bookEvent.Wait()
}