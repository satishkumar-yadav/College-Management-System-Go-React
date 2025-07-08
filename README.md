# 🎓 College Management System (React-Go Full Stack)

Full stack college portal using React + Go + MySQL + MongoDB.

## 🔧 Tech Stack

- 🧠 Frontend: React + Vite + Tailwind/CSS
- ⚙️ Backend: Go (Golang)
- 💾 Databases: MySQL (Main), MongoDB (Temporary Storage)
- 🔐 Auth: JWT + Email (SMTP)
- 📦 Dockerized later 

## 🗂️ Features

- Self Register (Faculty/Student) → Admin Approval
- Dynamic Login Dashboard
- Placement & Achievements
- Virtual College Tour
- Notices, Testimonials
- Admin Panel with Approve/Reject
- Email Notifications

## 📁 Project Structure

- college-management-system/
- ├── backend/cms/
- │ ├── main.go
- │ ├── go.mod
- │ ├── db/
- │ │ ├── mysql.go
- │ │ └── mongo.go
- │ ├── handlers/
- │ │ ├── auth.go
- │ │ ├── dashboard.go
- │ │ ├── notice.go
- │ │ └── admin.go
- │ ├── models/
- │ │ └── user.go
- │ ├── middleware/
- │ │ └── jwtMiddleware.go
- │ └── utils/
- │ └── email.go
- ├── frontend/
- │ ├── public/
- │ │ └── index.html
- │ ├── src/
- │ │ ├── App.jsx
- │ │ ├── main.jsx
- │ │ ├── index.css
- │ │ ├── components/
- │ │ │ ├── Navbar.jsx
- │ │ │ ├── Carousel.jsx
- │ │ │ ├── LoginRegisterModal.jsx
- │ │ │ ├── NoticeSidebar.jsx
- │ │ │ ├── AboutCollege.jsx
- │ │ │ ├── PlacementAchievements.jsx
- │ │ │ ├── VirtualTour.jsx
- │ │ │ ├── Dashboard.jsx
- │ │ │ └── AdminPanel.jsx
- │ │ └── styles/
- │ │ └── all-section-css-files.css
- │ ├── Dockerfile
- │ ├── package.json
- │ └── vite.config.js
- ├── docker-compose.yml
- ├── .gitignore
- ├── .env
- ├── LICENSE.txt
- ├── Go Dependencies.txt
- └── README.md

## 🚀 Getting Started

```bash
# 1. Rename .env.example to .env
cp .env.example .env

# 2. Run Docker containers
docker-compose up --build
```
