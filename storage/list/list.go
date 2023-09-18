package list

import "fmt"

type List struct {
	len       int64
	firstNode *node
}

// NewList создает новый список
func NewList() (l *List) {
	return &List{len: 0, firstNode: nil}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value int64) (id int64) {
	newNode := &node{value: value}

	if l.firstNode == nil {
		l.firstNode = newNode
		l.len++
		l.firstNode.index = 0
		return 0
	}

	nn := l.firstNode
	for ; nn.next != nil; nn = nn.next {
	}

	newNode.index = nn.index + 1
	nn.next = newNode

	l.len++

	return newNode.index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	if l.firstNode == nil {
		return
	}

	if id == l.firstNode.index {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}

	for nn := l.firstNode; nn.next != nil; nn = nn.next {
		if nn.next.index == id {
			nn.next = nn.next.next
			l.len--
			return
		}
	}
	return
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value int64) {
	if l.firstNode == nil {
		return
	}

	if value == l.firstNode.value {
		l.firstNode = l.firstNode.next
		l.len--
		return
	}

	for nn := l.firstNode; nn.next != nil; nn = nn.next {
		if nn.next.value == value {
			nn.next = nn.next.next
			l.len--
			return
		}
	}
	return
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value int64) {
	if l.firstNode == nil {
		return
	}

	if value == l.firstNode.value {
		l.firstNode = l.firstNode.next
		l.len--
	}

	for nn := l.firstNode; nn.next != nil; {
		if nn.next.value == value {
			nn.next = nn.next.next
			l.len--
			if nn.next == nil {
				break
			}
		} else {
			nn = nn.next
		}
	}
	return
}

// GetByIndex возвращает значение элемента по индексу.
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value int64, ok bool) {

	if l.firstNode == nil {
		return 0, false
	}

	for nn := l.firstNode; nn != nil; nn = nn.next {
		if nn.index == id {
			return nn.value, true
		}
	}
	return
}

// GetByValue возвращает индекс первого найденного элемента по значению.
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value int64) (id int64, ok bool) {
	if l.firstNode == nil {
		return 0, false
	}

	for nn := l.firstNode; nn != nil; nn = nn.next {
		if nn.value == value {
			return nn.index, true
		}
	}
	return
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value int64) (ids []int64, ok bool) {
	if l.firstNode == nil {
		return nil, false
	}

	for nn := l.firstNode; nn != nil; nn = nn.next {
		if nn.value == value {
			ids = append(ids, nn.index)
		}
	}

	if ids == nil {
		return nil, false
	}
	return ids, true
}

// GetAll возвращает все элементы списка
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []int64, ok bool) {
	if l.firstNode == nil {
		return nil, false
	}

	for nn := l.firstNode; nn != nil; nn = nn.next {
		values = append(values, nn.value)
	}

	return values, true
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	if l.firstNode == nil {
		fmt.Println("Пусто")
		return
	}
	for n := l.firstNode; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
