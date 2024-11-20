import React, { useEffect, useState } from 'react';
import { Grid, Paper, Typography, Alert } from '@mui/material';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
} from 'chart.js';
import { Bar } from 'react-chartjs-2';
import { api } from '../services/api';

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);

interface DashboardStats {
  totalSpaces: number;
  availableSpaces: number;
  occupiedSpaces: number;
  hourlyData: {
    hour: string;
    occupied: number;
  }[];
}

const Dashboard = () => {
  const [stats, setStats] = useState<DashboardStats | null>(null);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const data = await api.getMonitoringData();
        setStats(data);
        setError(null);
      } catch (err) {
        setError('Failed to load dashboard data');
        console.error('Dashboard data error:', err);
      }
    };
    fetchData();
  }, []);

  const chartData = {
    labels: stats?.hourlyData?.map(d => d.hour) || [],
    datasets: [
      {
        label: 'Occupied Spaces',
        data: stats?.hourlyData?.map(d => d.occupied) || [],
        backgroundColor: 'rgba(53, 162, 235, 0.5)',
        borderColor: 'rgb(53, 162, 235)',
        borderWidth: 1,
      },
    ],
  };

  const chartOptions = {
    responsive: true,
    plugins: {
      legend: {
        position: 'top' as const,
      },
      title: {
        display: true,
        text: 'Hourly Parking Occupancy',
      },
    },
    scales: {
      y: {
        beginAtZero: true,
      },
    },
  };

  if (error) {
    return <Alert severity="error">{error}</Alert>;
  }

  return (
    <Grid container spacing={3}>
      <Grid item xs={12}>
        <Typography variant="h4" gutterBottom>
          Parking Space Dashboard
        </Typography>
      </Grid>
      
      {/* Stats Cards */}
      <Grid item xs={12} md={4}>
        <Paper sx={{ p: 2, bgcolor: '#e3f2fd' }}>
          <Typography variant="h6" color="primary">Total Spaces</Typography>
          <Typography variant="h3">{stats?.totalSpaces ?? 0}</Typography>
        </Paper>
      </Grid>

      <Grid item xs={12} md={4}>
        <Paper sx={{ p: 2, bgcolor: '#e8f5e9' }}>
          <Typography variant="h6" color="success.main">Available Spaces</Typography>
          <Typography variant="h3">{stats?.availableSpaces ?? 0}</Typography>
        </Paper>
      </Grid>

      <Grid item xs={12} md={4}>
        <Paper sx={{ p: 2, bgcolor: '#ffebee' }}>
          <Typography variant="h6" color="error">Occupied Spaces</Typography>
          <Typography variant="h3">{stats?.occupiedSpaces ?? 0}</Typography>
        </Paper>
      </Grid>

      {/* Chart */}
      <Grid item xs={12}>
        <Paper sx={{ p: 3 }}>
          <Bar options={chartOptions} data={chartData} />
        </Paper>
      </Grid>
    </Grid>
  );
};

export default Dashboard;