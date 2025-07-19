# 🎓 College Management System (React-Go Full Stack)

Full stack college portal using React + Go + MySQL + MongoDB.

A full-featured **College Management System** built using Java Swing (Desktop App) and JSP/Servlet (Web App) with MySQL database and AWS deployment.

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

## 🔑 Key Features
- 🔐 Student, Faculty, Admin login system
- 🧾 Attendance & Result Management
- 💳 Fee Payment and Receipt Generation
- 📊 Personalized Dashboards
- 🌐 Online Deployment on AWS EC2
- 🗂️ Session Timeout & Role-based Access

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

## 🖼️ Screenshots
> _Add screenshots of your app UI_

## 🚀 How to Run
1. Clone the repository
2. Import Java Swing frontend in IntelliJ/Eclipse
3. Configure backend using Apache Tomcat for JSP/Servlet
4. Import SQL schema into MySQL
5. Update JDBC DB credentials in `.java` files

## 🚀 Getting Started

```bash
# 1. Rename .env.example to .env
cp .env.example .env

# 2. Run Docker containers
docker-compose up --build
```
## 🌍 Deployed Link
> http://your-aws-ec2-link

## 👨‍💻 Author
- **Your Name** – [LinkedIn](https://www.linkedin.com/in/satishkumar-yadav) | [GitHub](https://github.com/satishkumar-yadav)

## 📜 License
MIT License
