# Eventor
Eventor aims to provide an utility library to support event sourcing.

For further understanding of event sourcing [this publication](http://microservices.io/patterns/data/event-sourcing.html) by Chris Richardson is a good starting point.

## Goals
The library aims to:
- [x] Define event listener functions that updates business entities
- [x] Abstract communication with the event store
    - [x] Support for kafka as event store (not decoupled from kafka)
- [ ] Replay events on startup to build actual status of the business entity
- [x] Fire events on changes to the entities
