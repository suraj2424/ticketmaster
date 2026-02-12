# High Concurrency Event Booking System
This project describes the technical difficulties of building a system that can handle a massive number of users trying to book tickets (or access limited resources) simultaneously, without over-selling, crashing, or providing a poor user experience.

**Purpose:** Understand `LOCKING` of system.\
**Use Case:** Preventing over-selling of limited inventory in real-time.\

**Core Challenges:**
- **Distributed Locking:** Using Redis to lock a seat while a user is in the "checkout" phase so no one else can touch it.
- **Atomicity:** Ensuring that "Reserve Seat" → "Process Payment" → "Confirm Seat" happens as a single atomic unit or rolls back correctly.
- **Load Shedding:** Implementing rate-limiting so the server doesn't crash during a traffic spike.

## TECH STACK
> GO, Redis, Database (Any), Stripe(Future updates)

## RESULTS
API Documentation (Swagger/OpenAPI) & Load Testing Reports

**To test the system**
- We can use a tool like `k6` or `GO`'s own benchmark tools to simulate 5,000 users hitting your api at once.

