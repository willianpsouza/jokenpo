# 🕹️ Korean-Style Jokenpo CLI in Golang

![Jokenpo Korean Style](https://pm1.aminoapps.com/6638/1aeb77cb1a11ed25a7cae88ac60cccd123dafc10_hq.jpg)

## 🎮 About the Game

This is a **Command-Line Interface (CLI) Jokenpo** game written in **Golang**, with **data storage in SQLite**. The project is fully modularized, with each module handling different functionalities, including **randomness control based on prime numbers** and **secure input masking**.

### 🛠️ Features:
✅ Play against **historical computers** (IBM 386, Apple II, ZX Spectrum, etc.)  
✅ Play against **another human**, with masked input to prevent cheating  
✅ Store results in **SQLite**, using a **rudimentary blockchain mechanism**  
✅ List your past **game results**

---

## 🕹️ How to Play

### **Option 1 - Play Against the Computer**
- Enter your **name**
- Choose:  
  1️⃣ Rock 🪨  
  2️⃣ Paper 📄  
  3️⃣ Scissors ✂️
- The **computer picks a historical computer** as your opponent (IBM 386, Apple II, etc.)
- The result is **stored in the SQLite database**

### **Option 2 - Play Against Another Human**
- **Both players enter their names**
- **Masked input** prevents the second player from seeing the first player's choice
- The result is **stored in the SQLite database**

### **Option 3 - View Your Game Results**
- Enter your **name**
- See a list of your past **matches and winners**

---

## 🔐 Blockchain-Like Database Integrity

To prevent **game tampering**, the **database stores a cryptographic hash** for each game result:
- The **player names**
- Their **choices**
- A **timestamp**
- A **hash of the last saved entry**

This creates a **basic blockchain mechanism**. In the future, a function can **validate if any game record was altered**.

---

## 🏗️ Technical challenges

- The prime-based random module is simple but serves the purpose of challenge and game play, the Harmony VRF mode given the sophistication of the Harmony VRF module, I believe its best use is in a blockchain where the public verification characteristics of the random value could be validated by anyone.
- Tests are full implemented.
- One Makefile has been provided, to make easy run in container. `$make all` will perform like charm.
- Jokenpo.yaml file create artifact to be used inside simple container. 
```dockerfile
FROM alpine:latest
WORKDIR /app
COPY . .
ENV GOMAXPROCS=4
ENTRYPOINT ["/app/jokenpo"]
```
- At race conditions tests randomize tools show weakness, and context mutex has been implemented, to avoid race conditions.

```bash
go clean -testcache
CGO_ENABLED=1 go test -race  ./...
?   	jokenpo.provengo.io	[no test files]
?   	jokenpo.provengo.io	[no test files]
ok  	jokenpo.provengo.io/internal/dbhandler	1.022s
ok  	jokenpo.provengo.io/internal/encrypt	2.062s
ok  	jokenpo.provengo.io/internal/game	1.195s
ok  	jokenpo.provengo.io/internal/randomize	1.009s
ok  	jokenpo.provengo.io/internal/setup	1.010s
ok  	jokenpo.provengo.io/internal/utils	1.010s
```

### Why Multiply Prime Numbers for Randomness?

The idea behind multiplying prime numbers to generate random numbers comes from mathematical and cryptographic principles. 

### Here are the key reasons why this approach is used:

- Chaotic Distribution:
When prime numbers are multiplied, the resulting product appears to have a more chaotic distribution. Since prime numbers do not follow a simple pattern, their multiplication helps in generating numbers that seem random.

- Hard to Reverse (Factorization Problem):
In cryptography, the multiplication of two large prime numbers is considered a one-way function, meaning it is computationally difficult to reverse (i.e., factorize the product to find the original primes). This is the foundation of RSA encryption.

- Modular Arithmetic & Cycles:
When using modular arithmetic with prime numbers—such as —the resulting values tend to have properties that make them harder to predict, making this useful for pseudo-random number generation (PRNGs).

- Use in Pseudo-Random Generators:
Prime numbers play a crucial role in PRNGs. The Blum Blum Shub algorithm, for example, relies on the multiplication of two large prime numbers in modular arithmetic to generate secure random sequences.

## 📜 References: 
- Menezes, A., van Oorschot, P., & Vanstone, S. (1996). Handbook of Applied Cryptography. CRC Press.
- Koblitz, N. (1994). A Course in Number Theory and Cryptography. Springer.
- Blum, L., Blum, M., & Shub, M. (1986). A Simple Unpredictable Pseudo-Random Number Generator. SIAM Journal on Computing.

---
## Dependency
### **Golang Last Version **
🛠️ Installing Go

To run the Go example above, you need to have Go installed on your system. Follow the official installation instructions at: [Go Installation Guide](https://go.dev/doc/install)

## ⚡ Installation & Usage

### **1️⃣ Clone the Repository**
```sh
git clone https://github.com/willianpsouza/jokenpo.git
cd jokenpo
go run main.go
