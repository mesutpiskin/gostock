import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import ListItemAvatar from '@material-ui/core/ListItemAvatar';
import Avatar from '@material-ui/core/Avatar';
import LocalOfferIcon from '@material-ui/icons/LocalOffer';
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

export default function Dividers(props) {
  const classes = useStyles();
  console.log(props.fundModel.name);
  return (
    <List className={classes.root}>
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <LocalOfferIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Güncel Fiyat" secondary={"₺" + props.fundModel.price} />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <CallMissedOutgoingIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Günlük Getiri" secondary={props.fundModel.dailyreturn} />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <MoneyIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Toplam Pay" secondary={props.fundModel.pcs} />
      </ListItem>
      <Divider variant="inset" component="li" />
      <ListItem>
        <ListItemAvatar>
          <Avatar>
            <PeopleAltIcon color="primary" />
          </Avatar>
        </ListItemAvatar>
        <ListItemText primary="Fon Büyüklüğü" secondary={"₺" + props.fundModel.totalvalue} />
      </ListItem>
    </List>
  );
}