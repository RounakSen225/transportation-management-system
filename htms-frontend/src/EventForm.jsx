import React, { useState } from 'react';
import axios from 'axios';

const EventForm = ({ onEventCreated }) => {
  const [clientID, setClientID] = useState('');
  const [date, setDate] = useState('');
  const [cost, setCost] = useState('');
  const [serviceProvider, setServiceProvider] = useState('');
  const [datetime, setDatetime] = useState('');
  

  const handleSubmit = async (e) => {
    e.preventDefault();
    const eventData = {
      clientID: parseInt(clientID, 10),
      cost: parseFloat(cost),
      serviceProvider
    };

    console.log(eventData);

    try {
      const response = await axios.post('http://localhost:8080/events', eventData, {
        headers: {
            'Content-Type': 'application/json'
        }
    });
      onEventCreated(response.data);
      alert('Event recorded successfully');
      checkForAlerts(clientID);
    } catch (error) {
      console.error('Failed to record event:', error);
      alert('Failed to record event');
    }
  };

  const checkForAlerts = async (clientId) => {
    try {
      const response = await axios.get(`http://localhost:8080/clients/${clientId}/alerts`);
      // Handle the response if alerts are present
      if (response.data.length > 0) {
        alert('Alert triggered! ' + response.data.map(alert => alert.message).join(", "));
      }
    } catch (error) {
      console.error('Failed to check for alerts:', error);
    }
  };

  const handleDateTimeChange = (e) => {
    const value = e.target.value;

    // Assuming the backend expects UTC time and the input provides local time
    // Format the datetime as ISO string without 'Z'
    // If seconds are needed, you would handle them here (most browsers handle seconds if included in the input)
    const formattedDateTime = value.replace('Z', '.000');
    setDatetime(formattedDateTime);
 };


  return (
    <form onSubmit={handleSubmit}>
      <h2>Record Transportation Event</h2>
      <div>
        <label>Client ID:</label>
        <input
          type="number"
          value={clientID}
          onChange={(e) => setClientID(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Cost:</label>
        <input
          type="number"
          value={cost}
          onChange={(e) => setCost(e.target.value)}
          required
        />
      </div>
      <div>
        <label>Service Provider:</label>
        <input
          type="text"
          value={serviceProvider}
          onChange={(e) => setServiceProvider(e.target.value)}
          required
        />
      </div>
      <button type="submit">Submit</button>
    </form>
  );
};

export default EventForm;
