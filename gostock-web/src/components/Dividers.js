import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import ImageIcon from '@material-ui/icons/Image';
import LocalOfferIcon from '@material-ui/icons/LocalOffer';
import BeachAccessIcon from '@material-ui/icons/BeachAccess';
import CallMissedOutgoingIcon from '@material-ui/icons/CallMissedOutgoing';
import Divider from '@material-ui/core/Divider';
import PeopleAltIcon from '@material-ui/icons/PeopleAlt';
import MoneyIcon from '@material-ui/icons/Money';

const useStyles = makeStyles((theme) => ({
  root: {
    width: '100%',
    backgroundColor: theme.palette.background.paper,
  },
}));

export default function Dividers() {
  const classes = useStyles();

  return (
    <List className={classes.root}>
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <LocalOfferIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Son Fiyat (TL)" secondary="13,142674" />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <CallMissedOutgoingIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Günlük Getiri (%)" secondary="13,142674" />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <MoneyIcon color="primary"/>
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Pay (Adet)" secondary="13,142674" />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <PeopleAltIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Fon Toplam Değer (TL)" secondary="13,142674" />
      </ListItem>
    </List>
  );
}