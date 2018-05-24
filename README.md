# Taskar App
The contextual awareness project, which supports the autonomous
wheelchair project, uses this service as a receptor for sensor data.
the microcontroller connects to this app periodically and sends the
readings from all the peripheral sensor devices.

This app is written in Go, and it starts a server to receive the data.
The app will also host a UI webapp for visualizing data. Some back-end
data service or structure is used to persist the data.
