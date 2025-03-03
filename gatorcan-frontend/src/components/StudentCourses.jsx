import React, { useEffect, useState } from "react";
import StudentNavbar from "./StudentNavbar";

function StudentCourses() {
  const [allCourses, setAllCourses] = useState([]);
  const [enrolledCourses, setEnrolledCourses] = useState([]);
  const [loadingAllCourses, setLoadingAllCourses] = useState(true);
  const [loadingEnrolledCourses, setLoadingEnrolledCourses] = useState(true);

  useEffect(() => {
    loadCourses();
    loadEnrolledCourses();
  }, []);

  const loadCourses = async () => {
    setLoadingAllCourses(true);
    const courses = [
      {
        id: 1,
        name: "Introduction to Python",
        description: "Learn the basics of Python programming.",
        created_at: "2024-07-11T05:30:46Z",
        updated_at: "2024-07-22T05:30:46Z",
      },
      {
        id: 2,
        name: "Advanced Java",
        description: "Master Java programming with advanced topics.",
        created_at: "2025-01-19T05:30:46Z",
        updated_at: "2025-01-19T05:30:46Z",
      },
      {
        id: 3,
        name: "Machine Learning Basics",
        description: "Understand the fundamentals of machine learning.",
        created_at: "2024-11-23T05:30:46Z",
        updated_at: "2024-12-21T05:30:46Z",
      },
      {
        id: 4,
        name: "Deep Learning Fundamentals",
        description: "Dive into deep learning concepts and frameworks.",
        created_at: "2024-04-05T05:30:46Z",
        updated_at: "2024-04-27T05:30:46Z",
      },
      {
        id: 5,
        name: "Data Structures & Algorithms",
        description: "Learn essential data structures and algorithms.",
        created_at: "2024-08-06T05:30:46Z",
        updated_at: "2024-08-07T05:30:46Z",
      },
      {
        id: 6,
        name: "Cloud Computing with AWS",
        description: "Get hands-on experience with AWS cloud computing.",
        created_at: "2024-06-14T05:30:46Z",
        updated_at: "2024-07-10T05:30:46Z",
      },
      {
        id: 7,
        name: "Full Stack Web Development",
        description: "Build full-stack web applications from scratch.",
        created_at: "2024-10-01T05:30:46Z",
        updated_at: "2024-10-21T05:30:46Z",
      },
      {
        id: 8,
        name: "Big Data Analytics",
        description: "Analyze large datasets using modern big data tools.",
        created_at: "2024-05-22T05:30:46Z",
        updated_at: "2024-06-15T05:30:46Z",
      },
      {
        id: 9,
        name: "Cybersecurity Essentials",
        description:
          "Learn the principles of cybersecurity and best practices.",
        created_at: "2024-12-12T05:30:46Z",
        updated_at: "2025-01-10T05:30:46Z",
      },
      {
        id: 10,
        name: "React for Beginners",
        description: "Start building dynamic web applications with React.",
        created_at: "2024-09-15T05:30:46Z",
        updated_at: "2024-09-29T05:30:46Z",
      },
    ];
    setAllCourses(courses || []);
    setLoadingAllCourses(false);
  };

  const loadEnrolledCourses = async () => {
    setLoadingEnrolledCourses(true);
    const courses = [
      
    ];
    setEnrolledCourses(courses);
    setLoadingEnrolledCourses(false);
  };

  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "120px", padding: "20px" }}>
        <h1>Courses</h1>
        <hr />

        <h4>Enrolled Courses</h4>
        <hr />
        {loadingEnrolledCourses ? (
          <p>Loading enrolled courses...</p>
        ) : enrolledCourses.length === 0 ? (
          <p>No enrolled courses</p>
        ) : (
          <div className="grid-container">
            {enrolledCourses.map((course) => (
              <CourseCard key={course.id} course={course} isEnrolled />
            ))}
          </div>
        )}

        <br />

        <h4>All Courses</h4>
        <hr />
        {loadingAllCourses ? (
          <p>Loading courses...</p>
        ) : allCourses.length === 0 ? (
          <p>No course available</p>
        ) : (
          <div className="grid-container">
            {allCourses.map((course) => (
              <CourseCard
                key={course.id}
                course={course}
              />
            ))}
          </div>
        )}
      </div>
      <style>
        {`
    .grid-container {
      display: grid;
      grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
      gap: 20px;
      padding: 20px;
    }

    @media (max-width: 900px) {
      .grid-container {
        grid-template-columns: repeat(2, minmax(300px, 1fr));
      }
    }

    @media (max-width: 600px) {
      .grid-container {
        grid-template-columns: repeat(1, minmax(300px, 1fr));
      }
    }
  `}
      </style>
    </>
  );
}

const CourseCard = ({ course, isEnrolled = false }) => {
  return (
    <div
      style={{
        border: "1px solid #ddd",
        borderRadius: "8px",
        padding: "15px",
        boxShadow: "2px 2px 10px rgba(0,0,0,0.1)",
        display: "flex",
        flexDirection: "column",
        justifyContent: "space-between",
        minHeight: "200px",
        position: "relative",
      }}
    >
      <h3 style={{ textAlign: "center" }}>{course.name}</h3>
      <p style={{ textAlign: "left", marginTop: "5px" }}>
        {course.description}
      </p>

      <div
        style={{
          display: "flex",
          justifyContent: "space-between",
          fontSize: "14px",
          marginTop: "10px",
          color: "#555",
        }}
      >
        <span>
          <strong>Created:</strong>{" "}
          {new Date(course.created_at).toLocaleDateString()}
        </span>
        <span>
          <strong>Updated:</strong>{" "}
          {new Date(course.updated_at).toLocaleDateString()}
        </span>
      </div>

      {isEnrolled ? (
        <>
          <p style={{ textAlign: "left", marginTop: "10px" }}>
            <strong>Instructor:</strong> {course.instructorName} (
            {course.instructorEmail})
          </p>
          <p
            style={{
              color: "green",
              fontWeight: "bold",
              textAlign: "center",
              marginTop: "10px",
            }}
          >
            Enrolled
          </p>
        </>
      ) : (
          <div></div>
      )}
    </div>
  );
};

export default StudentCourses;
