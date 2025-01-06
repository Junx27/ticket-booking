## Booking Documentation

In this booking transaction, the user will make a booking with the initial booking status of "pending", followed by an update to the seat availability. If the user decides to cancel the booking (i.e., the booking status changes to "cancel"), the seat availability will be restored based on the selected booking data. The same applies when deleting a booking by ID. If the booking status is "pending" or "payment", the seat will be restored. However, if the booking status is "payment" and the booking is canceled, it is important for the admin to take action on the canceled data. The admin will update the seat count based on the canceled booking and initiate a refund process for the customer who made the ticket reservation.

- cancellation without charge

```json
{
  "booking_status": "cancel"
}
```

before

```json
{
  "available_seats": [1, 2]
}
```

after

```json
{
  "available_seats": [1, 2, 3]
}
```

- cancellation after payment

```json
{
  "booking_status": "cancel"
}
```

before

```json
{
  "available_seats": [1, 2]
}
```

after

```json
{
  "available_seats": [1, 2, 3]
}
```

- cancellation status pending

```json
{
  "booking_status": "cancel"
}
```

before

```json
{
  "available_seats": [1, 2]
}
```

after

```json
{
  "available_seats": [1, 2, 3]
}
```

- deleting status pending

```json
{
  "booking_status": "cancel"
}
```

before

```json
{
  "available_seats": [1, 2]
}
```

after

```json
{
  "available_seats": [1, 2, 3]
}
```

[Back to Main Documentation](../README.md)
