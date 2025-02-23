package application

import (
	"sync"
)

type BookEvent struct {
	mu        sync.Mutex
	listeners []chan struct{}
}

var bookEvent = &BookEvent{}

// Notifica a todos los listeners y limpia la lista
func (e *BookEvent) NewNotifyBook() {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Notifica a todos los listeners sin bloquear
	for _, ch := range e.listeners {
		select {
		case ch <- struct{}{}: // Intenta enviar sin bloquear
		default:               // Si ya hay un valor en el buffer, no bloquea
		}
	}

	// Limpia la lista de listeners
	e.listeners = nil
}

// Registra un listener y devuelve un canal que será notificado cuando haya cambios
func (e *BookEvent) Wait() <-chan struct{} {
	e.mu.Lock()
	defer e.mu.Unlock()

	// Crea un canal CON BUFFER para evitar bloqueos
	ch := make(chan struct{}, 1)

	// Agrega el canal a la lista de listeners
	e.listeners = append(e.listeners, ch)

	return ch
}

// Función global para notificar actualizaciones de libros
func NotifyBookUpdate() {
	bookEvent.NewNotifyBook()
}

// Función global para esperar actualizaciones de libros
func WaitForBookUpdate() <-chan struct{} {
	return bookEvent.Wait()
}
