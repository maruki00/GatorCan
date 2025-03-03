import axios from "axios";
import { fetchAllCourses } from "./CourseService";

jest.mock("axios");

describe("CourseService - fetchAllCourses", () => {
  test("successfully fetches all courses", async () => {
    const mockCourses = {
      courses: [
        { id: 1, name: "Data Science", description: "Intro to Data Science" },
        { id: 2, name: "Software Engineering", description: "SE Principles" },
      ],
    };

    axios.get.mockResolvedValue({ data: mockCourses });

    const result = await fetchAllCourses();

    expect(result).toEqual(mockCourses);
    expect(axios.get).toHaveBeenCalledWith(
      "http://localhost:8080/courses/?page=1&pageSize=10",
      expect.objectContaining({
        headers: expect.objectContaining({
          Authorization: expect.stringContaining("Bearer "),
        }),
      })
    );
  });
});
