import React from 'react';
import axios from 'axios';
import { useTable } from 'react-table';
import { Column } from 'react-table';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
import EventForm from './EventForm';

interface Client {
  id: number;
  name: string;
  address: string;
  budget: number;
}

interface Event {
  id: number;
  clientID: number;
  date: string;
  cost: number;
  serviceProvider: string;
}

interface Alert {
    id: number;
    clientID: number;
    message: string;
    triggered: boolean;
    value: number;
  }  

const App: React.FC = () => {
  const [alerts, setAlerts] = React.useState<Alert[]>([]);
  const [selectedClientId, setSelectedClientId] = React.useState<number | null>(null);
  const [activeTab, setActiveTab] = React.useState<'clients' | 'events'>('clients');
  const [clients, setClients] = React.useState<Client[]>([]);
  const [events, setEvents] = React.useState<Event[]>([]);
  const [loading, setLoading] = React.useState<boolean>(true);
  
  const handleEventCreated = (eventData: Event) => {
    console.log('New event created:', eventData);
    // Additional handling can be implemented here if needed
  };
 
  const clientsColumns : Column<Client>[] = React.useMemo(
    () => [
        { Header: 'Name', accessor: 'name' as keyof Client},
        { Header: 'Address', accessor: 'address' as keyof Client},
        { Header: 'Budget', accessor: 'budget' as keyof Client},
    ],
    []
);

const eventsColumns : Column<Event>[] = React.useMemo(
    () => [
        { Header: 'Date', accessor: 'date' as keyof Event},
        { Header: 'Cost', accessor: 'cost' as keyof Event},
        { Header: 'Service Provider', accessor: 'serviceProvider' as keyof Event},
    ],
    []
);


// Function to fetch alerts for a client
const fetchAlerts = async (clientId: number) => {
  try {
    const response = await axios.get<Alert[]>(`http://localhost:8080/clients/${clientId}/alerts`);
    setAlerts(response.data);
  } catch (error) {
    console.error('Error fetching alerts:', error);
    // Handle error appropriately in your actual UI
  }
};

// You would call fetchAlerts when a user selects a client, for example


// Use react-table for tabular data
const clientsTableInstance = useTable<Client>({ columns: clientsColumns, data: clients });
const eventsTableInstance = useTable<Event>({ columns: eventsColumns, data: events });

// Function to fetch data from backend
const fetchData = async () => {
    try {
        const clientsResponse = await axios.get<Client[]>('http://localhost:8080/clients');
        const eventsResponse = await axios.get<Event[]>('http://localhost:8080/events');
        setClients(clientsResponse.data);
        setEvents(eventsResponse.data);
    } catch (error) {
        console.error('Error fetching data:', error);
    } finally {
        setLoading(false);
    }
};

// Effect to run the fetchData function on component mount
React.useEffect(() => {
    fetchData();
}, []);

React.useEffect(() => {
    if (selectedClientId) {
      fetchAlerts(selectedClientId);
    }
  }, [selectedClientId]);

if (loading) {
    return <div>Loading...</div>;
}

const handleTabClick = (tab: 'clients' | 'events') => {
    setActiveTab(tab);
};


// Create a basic bar graph for client budgets using Recharts
const budgetData = clients.map(client => ({ name: client.name, budget: client.budget }));

return (
    <div className="dashboard">
        <h1>Transportation Management Dashboard</h1>
        <EventForm onEventCreated={handleEventCreated} />
        <div>
            <button onClick={() => handleTabClick('clients')}
            className={activeTab === 'clients' ? 'active' : ''}>Clients</button>
            <button onClick={() => handleTabClick('events')}
            className={activeTab === 'events' ? 'active' : ''}>Events</button>
        </div>
        <div className="content">
            {activeTab === 'clients' ? (
            <div className="section">
            <h2>Clients</h2>
            <ClientsTable {...clientsTableInstance} />
            <BarChart width={600} height={300} data={budgetData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="name" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Bar dataKey="budget" fill="#8884d8" />
            </BarChart>
            </div>
            ) : ( 
            <div className='section'>    
            <EventsTable {...eventsTableInstance} />
            <BarChart width={600} height={300} data={budgetData}>
                <CartesianGrid strokeDasharray="3 3" />
                <XAxis dataKey="date" />
                <YAxis />
                <Tooltip />
                <Legend />
                <Bar dataKey="cost" fill="#8884d8" />
            </BarChart>
            </div>
            )}
        </div>
    <div className="alerts">
    <h3>Alerts</h3>
    {alerts.length > 0 ? (
        <ul>
        {alerts.map((alert) => (
            <li key={alert.id} className={alert.triggered ? 'alert-triggered' : ''}>
            {alert.message} - Value: ${alert.value}
            </li>
        ))}
        </ul>
    ) : (
        <p>No alerts to display</p>
    )}
    </div>
    </div>
    );
    
}

const ClientsTable: React.FC<any> = ({ getTableProps, getTableBodyProps, headerGroups, rows, prepareRow }) => (
    <table {...getTableProps()}>
        <thead>
            {headerGroups.map((headerGroup: any) => (
                <tr {...headerGroup.getHeaderGroupProps()}>
                    {headerGroup.headers.map((column: any) => (
                        <th {...column.getHeaderProps()}>{column.render('Header')}</th>
                    ))}
                </tr>
            ))}
        </thead>
        <tbody {...getTableBodyProps()}>
            {rows.map((row : any) => {
                prepareRow(row);
                return (
                    <tr {...row.getRowProps()}>
                        {row.cells.map((cell: any) => {
                            return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>;
                        })}
                    </tr>
                );
            })}
        </tbody>
    </table>
);

const EventsTable: React.FC<any> = ({ getTableProps, getTableBodyProps, headerGroups, rows, prepareRow }) => (
    <table {...getTableProps()}>
        <thead>
            {headerGroups.map((headerGroup: any) => (
                <tr {...headerGroup.getHeaderGroupProps()}>
                    {headerGroup.headers.map((column: any) => (
                        <th {...column.getHeaderProps()}>{column.render('Header')}</th>
                    ))}
                </tr>
            ))}
        </thead>
        <tbody {...getTableBodyProps()}>
            {rows.map((row: any) => {
                prepareRow(row);
                return (
                    <tr {...row.getRowProps()}>
                        {row.cells.map((cell: any) => {
                            return <td {...cell.getCellProps()}>{cell.render('Cell')}</td>;
                        })}
                    </tr>
                );
            })}
        </tbody>
    </table>
);

export default App;
