import axios from "axios";

const base_url = "http://localhost:8080/courses";

const getAuthHeader = () => {
  const refreshToken = localStorage.getItem("refreshToken");
  return {
    Authorization: `Bearer ${refreshToken}`,
    "Content-Type": "application/json",
  };
};

// Fetch all available courses
export const fetchAllCourses = async () => {
  try {
    const response = await axios.get(`${base_url}/?page=1&pageSize=10`, {
      headers: getAuthHeader(),
    });
    return response.data;
  } catch (error) {
    console.error("Error fetching all courses:", error);
    return null;
  }
};

// Fetch enrolled courses
export const fetchEnrolledCourses = async () => {
  try {
    const response = await axios.get(`${base_url}/enrolled`, {
      headers: getAuthHeader(),
    });

    if (response.data === null) return [];

    return response.data.map((course) => ({
      id: course.ID,
      name: course.Name,
      description: course.Description,
      created_at: course.StartDate,
      updated_at: course.EndDate,
      instructorName: course.InstructorName,
      instructorEmail: course.InstructorEmail,
    }));
  } catch (error) {
    console.error("Error fetching enrolled courses:", error);
    return [];
  }
};

// Enroll in a course
export const enrollInCourse = async (courseID) => {
  try {
    const response = await axios.post(
      `${base_url}/enroll`,
      { courseID },
      { headers: getAuthHeader() }
    );

    alert(`Successfully enrolled in course ID: ${courseID}`);
    return { success: true };
  } catch (error) {
    console.error("Error enrolling in course:", error);
    alert("Enrollment failed!");
    return { success: false };
  }
};

export default {
  fetchAllCourses,
  fetchEnrolledCourses,
  enrollInCourse,
};
