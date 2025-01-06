## Booking Documentation

In this booking transaction, the user will make a booking with the initial booking status of "pending", followed by an update to the seat availability. If the user decides to cancel the booking (i.e., the booking status changes to "cancel"), the seat availability will be restored based on the selected booking data. The same applies when deleting a booking by ID. If the booking status is "pending" or "payment", the seat will be restored. However, if the booking status is "payment" and the booking is canceled, it is important for the admin to take action on the canceled data. The admin will update the seat count based on the canceled booking and initiate a refund process for the customer who made the ticket reservation.

# Booking Transaction Workflow:

# 1. Initial Booking Status: "pending"

# - A user makes a booking with the initial status set to "pending."

# - Following this, the seat availability is updated accordingly.

# 2. Cancellation of Booking:

# - If the user cancels the booking (i.e., booking_status changes to "cancel"),

# the seat availability will be restored based on the selected booking.

# 3. Deleting Booking by ID:

# - When a booking is deleted by ID, if the booking_status is "pending" or "payment",

# the seat availability will be restored.

# - However, if the booking status is "payment" and the booking is canceled,

# the admin must take action to follow up on the canceled booking.

# - Admin's actions:

# - Admin will update the seat availability based on the canceled booking.

# - Admin will initiate a refund process for the customer who made the ticket reservation.

# The admin's role in handling bookings and cancellations is crucial to ensure

# that seat availability is properly managed and customers who cancel are refunded.

[Back to Main Documentation](../README.md)
