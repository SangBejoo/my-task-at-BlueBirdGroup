
import React from 'react';
import { useEffect, useState } from 'react';
import { Box, Grid, Paper, Typography } from '@mui/material';
import { api } from '../services/api';

interface MapPlace {
  id: number;
  name: string;
  status: 'available' | 'occupied';
  type: string;
}

const ParkingMap = () => {
  const [places, setPlaces] = useState<MapPlace[]>([]);

  useEffect(() => {
    const fetchPlaces = async () => {
      try {
        const data = await api.getAllMapPlaces();
        setPlaces(data);
      } catch (error) {
        console.error('Error fetching parking places:', error);
      }
    };
    fetchPlaces();
  }, []);

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Parking Map
      </Typography>
      <Grid container spacing={2}>
        {places.map((place) => (
          <Grid item xs={2} key={place.id}>
            <Paper
              sx={{
                p: 2,
                textAlign: 'center',
                bgcolor: place.status === 'available' ? '#4caf50' : '#f44336',
                color: 'white',
              }}
            >
              <Typography variant="body1">{place.name}</Typography>
              <Typography variant="caption">{place.type}</Typography>
            </Paper>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
};

export default ParkingMap;