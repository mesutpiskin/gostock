import React from 'react';
import Link from '@material-ui/core/Link';
import { makeStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';

function preventDefault(event) {
  event.preventDefault();
}

const useStyles = makeStyles({
  depositContext: {
    flex: 1,
  },
});

export default function FundProfile() {
  const classes = useStyles();

  return (
    <React.Fragment>
     
      <Typography component="p" variant="h4">
      ₺3,024.00
      </Typography>
      <Typography color="textSecondary" className={classes.depositContext}>
      {new Date().toLocaleDateString()}
      </Typography>
      <div>
        <Link color="primary" href="#" onClick={preventDefault}>
        
        </Link>
      </div>
    </React.Fragment>
  );
}
