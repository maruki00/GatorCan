import React from 'react';
import StudentNavbar from './StudentNavbar';
import {ScheduleXCalendar, useCalendarApp} from "@schedule-x/react"
import {createViewWeek, createViewMonthGrid} from "@schedule-x/calendar";
import {createEventsServicePlugin} from "@schedule-x/events-service";
import '@schedule-x/theme-default/dist/index.css';
import { useState, useEffect } from 'react';
import { createEventModalPlugin } from "@schedule-x/event-modal";
import { createCurrentTimePlugin } from "@schedule-x/current-time";

function formatDate(date) {
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, "0");
  const day = String(date.getDate()).padStart(2, "0");
  const hours = String(date.getHours()).padStart(2, "0");
  const minutes = String(date.getMinutes()).padStart(2, "0");

  return `${year}-${month}-${day} ${hours}:${minutes}`;
}

function fetchEvents() {
  const days = ["S", "M", "T", "W", "R", "F", "S"];

  var courses = [
    {
      id: "1",
      title: "Intro to Data Science",
      start: "2025-01-01",
      end: "2025-05-01",
      days: ["M", "W", "F"],
      startTime: "10:40",
      endTime: "11:30",
      description: "Instructor: Christan Grant",
    },
    {
      id: "2",
      title: "Software Engineering",
      start: "2025-01-01",
      end: "2025-05-01",
      days: ["M", "W", "F"],
      startTime: "15:05",
      endTime: "15:55",
      description: "Instructor: Alin Dobra",
    },
    {
      id: "3",
      title: "Programming Language Principles",
      start: "2025-01-01",
      end: "2025-05-01",
      days: ["M", "W", "F"],
      startTime: "14:10",
      endTime: "15:00",
      description: "Instructor: Alin Dobra",
    },
  ];

  function createEventsFromCourses(courses) {
    var final_events = [];
    courses.forEach((course) => {
      const startDate = new Date(course["start"] + "T00:00:00");
      const endDate = new Date(course["end"] + "T00:00:00");
      for (
        var date = startDate;
        date <= endDate;
        date.setDate(date.getDate() + 1)
      ) {
        // your day is here
        const dayOfWeek = date.getDay();
        var dayOfWeekShort = days[dayOfWeek];
        if (course["days"].includes(dayOfWeekShort)) {
          var startTime = course["startTime"].split(":")[0];
          var endTime = course["endTime"].split(":")[1];
          var currentStartTime = date;
          var currentEndTime = date;
          currentStartTime.setHours(startTime, endTime, 0, 0);
          currentEndTime.setHours(startTime, endTime, 0, 0);
          final_events.push({
            id: course["id"],
            title: course["title"],
            start: formatDate(currentStartTime),
            end: formatDate(currentEndTime),
            description: course["description"],
          });
        }
      }
    });
    return final_events;
  }

  return createEventsFromCourses(courses);
}

function StudentCalendar() {

  const eventsService = useState(() => createEventsServicePlugin())[0];

  var events = fetchEvents();

  const calendar = useCalendarApp({
    views: [createViewWeek(), createViewMonthGrid()],
    events: events,
    selectedDate: formatDate(new Date()).slice(0, 10),
    plugins: [
      createEventModalPlugin(),
      eventsService,
      createCurrentTimePlugin({
        fullWeekWidth: true,
      }),
    ],
  });

  useEffect(() => {
    // get all events
    // eventsService.getAll();
  }, []);

  return (
    <>
      <StudentNavbar />
      <div style={{ marginLeft: "120px" }}>
        <h1>Calendar</h1>
        <hr />
        <br></br>
        <ScheduleXCalendar calendarApp={calendar} />
      </div>
    </>
  );
}

export default StudentCalendar;