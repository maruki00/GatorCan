import React from 'react';
import StudentNavbar from './StudentNavbar';
import {ScheduleXCalendar, useCalendarApp} from "@schedule-x/react"
import {createViewWeek, createViewMonthGrid} from "@schedule-x/calendar";
import {createEventsServicePlugin} from "@schedule-x/events-service";
import '@schedule-x/theme-default/dist/index.css';
import { useState, useEffect } from 'react';
import { createEventModalPlugin } from "@schedule-x/event-modal";

function StudentCalendar() {

  const eventsService = useState(() => createEventsServicePlugin())[0];

  const calendar = useCalendarApp({
    views: [createViewWeek(), createViewMonthGrid()],
    events: [
      {
        id: 1,
        title: "My New Event",
        start: "2025-01-01 00:00",
        end: "2025-01-01 02:00",
        description: "My cool description"
      },
    ],
    selectedDate: "2025-01-01",
    plugins: [createEventModalPlugin(), eventsService],
  });

  useEffect(() => {
    // get all events
    eventsService.getAll();
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