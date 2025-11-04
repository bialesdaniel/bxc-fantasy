
import { Box, Typography } from '@mui/material';
import { Code, ConnectError } from '@connectrpc/connect';
import { useEffect, useState } from 'react';

import { type CheckRequest, type CheckResponse } from '../../gen/health/v1/health_pb';
import { apiClient } from '../../lib/api';

const AboutPage = () => {
  const [currentUptime, setCurrentUptime] = useState(0);

  useEffect(() => {
    const controller = new AbortController();
    const getUptime = async () => {
      try {
        const res: CheckResponse = await apiClient.check({} as CheckRequest, {
          signal: controller.signal,
        });
        setCurrentUptime(res.uptimeSeconds);
        console.log(res);
      } catch (e) {
        if (e instanceof ConnectError && e.code === Code.Canceled) {
          return;
        }
        console.error(e);
      }
    };

    getUptime();
    return () => {
      controller.abort();
    };
  }, []);

  return (
    <Box>
      <Typography variant='h2' component='div'>
        About Page
      </Typography>
      <Typography variant='body1' component='div'>
        <b>Current Uptime:</b> {currentUptime}
      </Typography>
    </Box>
  );
};

export default AboutPage;
