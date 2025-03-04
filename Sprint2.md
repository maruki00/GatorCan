# 🏆 Sprint 2 - GatorCan

## 📅 Duration: [02/11/2025] - [03/03/2025]

## Visual Demo Links
- [Sprint 2 Integrated Demo](https://drive.google.com/file/d/1D9-meydP8ja-mxD-ICXWkcSgtUs_fDTe/view?usp=drivesdk)

## 🎯 Goal
Build the course management system with admin-controlled enrollment approval and implement the weekly schedule feature. Ensure seamless integration between backend and frontend.

---

## 📌 User Stories & Assignments

### **🔹 Backend (Mohammad & Muthu)**

#### **1️⃣ Define Database Schema for Courses & Enrollments (Mohammad)**
- **Who:** Backend Developers
- **Why:** To store and manage course and enrollment data efficiently.
- **What:** Implement schema for courses and enrollments, ensuring relationships are properly set up.

#### **2️⃣ Fetch Available Courses API (Mohammad)**
- **Who:** All users
- **Why:** To allow users to view available courses.
- **What:** Implement `GET /courses` endpoint with pagination and error handling.

#### **3️⃣ Course Enrollment API with Admin Approval (Muthu)**
- **Who:** Students (Request Enrollment), Admin (Approve/Reject)
- **Why:** To manage course enrollments with an approval process.
- **What:** Implement `POST /courses/enroll` with approval workflow and admin notifications.

#### **4️⃣ Fetch Enrolled Courses API (Muthu)**
- **Who:** Enrolled Students
- **Why:** To allow students to view their enrolled courses.
- **What:** Implement `GET /courses/enrolled` to fetch only courses that the user is enrolled in.

---

### **🔹 Frontend (Navnit & Harsh)**

#### **5️⃣ Course Listing Page UI (Navnit)**
- **Who:** All users
- **Why:** To allow users to browse available and enrolled courses.
- **What:** Design and implement a page that fetches data from `GET /courses` and `GET /courses/enrolled` APIs.

#### **6️⃣ Course Enrollment Request Workflow (Navnit)**
- **Who:** Students (Request Enrollment)
- **Why:** To enable students to submit enrollment requests.
- **What:** Implement a button to trigger `POST /courses/enroll` and display status updates.

#### **7️⃣ Weekly Schedule UI (Harsh)**
- **Who:** Enrolled Students
- **Why:** To display class timings and instructors based on enrolled courses.
- **What:** Design a UI showing a structured weekly schedule with course details.

#### **8️⃣ Backend API Integration for Courses & Schedule (Harsh)**
- **Who:** Frontend Developers
- **Why:** To connect UI components with backend functionality.
- **What:** Implement API calls to `GET /courses` and `GET /courses/enrolled` to dynamically populate the UI.

---

## ⚙️ **Sprint 2 - Issues & Completion Status**
### **Planned Issues:**
- Define and implement database schema
- Develop course-related API endpoints
- Design and build frontend course management UI
- Implement admin-controlled enrollment approval system
- Develop and integrate weekly schedule UI

### **Successfully Completed:** ✅ All planned issues were completed.

---

## 🚀 Outcome
By the end of Sprint 2, we have:
- ✅ Everything from Sprint 1
- ✅ Database schema for courses and enrollments
- ✅ API endpoints for course listing, enrollment, and approval workflow
- ✅ Course listing and enrollment UI
- ✅ Weekly schedule UI displaying enrolled courses and instructors
- ✅ Backend-frontend API integration

---

## 💚 Notes & Discussions
- [ ] Optimize database queries for large-scale course enrollments
- [ ] Improve UI responsiveness for weekly schedule
- [ ] Discuss potential enhancements for real-time notifications
- [ ] Plan for next sprint (Messaging System & Notifications)

---

### 🔥 Sprint 2 Successfully Completed! 🚀

