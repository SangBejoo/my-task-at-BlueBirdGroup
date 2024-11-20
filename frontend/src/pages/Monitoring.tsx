
import React, { useEffect, useState } from 'react';
import { Box, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';
import { api } from '../services/api';

interface Transaction {
  id: number;
  fleetNumber: string;
  driverId: string;
  placeId: number;
  status: string;
  createdAt: string;
}

const Monitoring = () => {
  const [transactions, setTransactions] = useState<Transaction[]>([]);

  useEffect(() => {
    const fetchTransactions = async () => {
      try {
        const data = await api.getAllTransactions();
        setTransactions(data);
      } catch (error) {
        console.error('Error fetching transactions:', error);
      }
    };
    fetchTransactions();
  }, []);

  return (
    <Box>
      <Typography variant="h4" gutterBottom>
        Parking Monitoring
      </Typography>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Fleet Number</TableCell>
              <TableCell>Driver ID</TableCell>
              <TableCell>Place ID</TableCell>
              <TableCell>Status</TableCell>
              <TableCell>Time</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {transactions.map((trx) => (
              <TableRow key={trx.id}>
                <TableCell>{trx.fleetNumber}</TableCell>
                <TableCell>{trx.driverId}</TableCell>
                <TableCell>{trx.placeId}</TableCell>
                <TableCell>{trx.status}</TableCell>
                <TableCell>{new Date(trx.createdAt).toLocaleString()}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Box>
  );
};

export default Monitoring;