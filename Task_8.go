package main

import (
    "fmt"
    "time"
)

type Guest struct {
    Name string
}

type Room struct {
    Number   int
    TypeRoom string
    Rate     float64
    Occupied bool
}

type Reservation struct {
    Guest     *Guest
    CheckIn   time.Time
    CheckOut  time.Time
    Room      *Room
}

type Hotel struct {
    Rooms []*Room
}

func (h *Hotel) IsAvailable(number int) bool {
    for _, r := range h.Rooms {
        if r.Number == number && !r.Occupied {
            fmt.Printf("Комната номер %d доступна\n", number)
            return true
        }
    }
    fmt.Printf("Комната номер %d занята или не существует\n", number)
    return false
}

func (h *Hotel) Reserve(guest *Guest, number int, checkIn, checkOut time.Time) (*Reservation, error) {
    for _, r := range h.Rooms {
        if r.Number == number && !r.Occupied {
            r.Occupied = true
            reserv := &Reservation{Guest: guest, CheckIn: checkIn, CheckOut: checkOut, Room: r}
            fmt.Printf("Комната номер %d забронирована для гостя %s\n", number, guest.Name)
            return reserv, nil
        }
    }
    return nil, fmt.Errorf("комната номер %d недоступна", number)
}

func (res *Reservation) CalculateCost() float64 {
    days := res.CheckOut.Sub(res.CheckIn).Hours() / 24
    cost := days * res.Room.Rate
    fmt.Printf("Стоимость проживания: %.2f рублей\n", cost)
    return cost
}

func main() {
    hotel := Hotel{
        Rooms: []*Room{{Number: 105, TypeRoom: "Стандарт", Rate: 3500.0}},
    }

    guest := Guest{Name: "Семёнова Елена Викторовна"}
    in := time.Date(2025, 9, 15, 0, 0, 0, 0, time.UTC)
    out := time.Date(2025, 9, 20, 0, 0, 0, 0, time.UTC)

    hotel.IsAvailable(105)
    reservation, _ := hotel.Reserve(&guest, 105, in, out)
    reservation.CalculateCost()
}
