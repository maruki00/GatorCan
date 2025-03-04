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

#### **✅ Unit Tests and Functional tests for Backend (Mohamamd and Muthu)**
##### Each of the test file covers all related functionalities
##### Unit Tests
- courseController -> Mocking courseService
    tests the controller functions like getCourse, getEnrolledCourse etc while integrated with the mocked service.
- userController -> Mocking userService
    tests the controller functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked service.
- userService -> Mocking userRepository, courseRepository
    tests the user service functions for implemented business logic for the functions like getUserDetails, login, updateUserDetails etc while integrated with the mocked reporsitories.
- courseService -> Mocking userRepository, courseRepository
    tests the course service functions for implemented business logic for the functions like getCourse, getEnrolledCourse etc while integrated with the mocked reporsitories.
##### Functonal Tests
- get enrolled courses
- enroll in_course
- get courses
- role based_access
- user deletion
- user details
- user login
- user registration
- user update

** Tested using positive, Negative and edge test cases.

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

#### **✅ Unit Tests and Cypress Test for Frontend (Harsh and Navnit)**
Cypress tests:
- fetch and validate login page
- validate the components in the login page
- pass in creds and click on login button
- fetch and validate dashboard page

Unit Tests:
- Login:
    1. Check if username and password are rendered correctly 
    2. Check if we are able to change username and password correctly 
    3. Check if we get an error message on passing invalid credentials 
- AdminDashboard:
    1. Check if add user tool renders correctly 
- StudentCourses:
    1. Check if enrolled courses and all courses heading is rendered properly 
    2. Check if after fetch all courses API runs (mock), it loads the courses onto the courses tab
    3. Check if "No enrolled courses" text renders if there are no courses enrolled by the student 
- StudentNavbar:
    1. Check if all Navbar components are rendered such as Profile, Calendar etc 
- AdminDashboard:
    1. Check if mock add user API gives correct success or failure responses
- AuthService:
    1. Check if mock login API gives correct success or failure response, and if local storage is updated with refresh token correctly 
- CourseService:
    1. Check if fetch all courses API gives correct success or failure responses.
- UserNavigation:
    1. Check if all elements are rendered correctly 
    2. Check if mock add user API gives correct display message on success or failure

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
- ✅ Unit tests and Cypress tests for backend and frontend

---

## 💚 Notes & Discussions
- [ ] Optimize database queries for large-scale course enrollments
- [ ] Improve UI responsiveness for weekly schedule
- [ ] Discuss potential enhancements for real-time notifications
- [ ] Plan for next sprint (Messaging System & Notifications)

---

### 🔥 Sprint 2 Successfully Completed! 🚀

