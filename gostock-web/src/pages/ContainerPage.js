import React from 'react';
import FundPage from './FundPage';
import Box from '@material-ui/core/Box';
import Typography from '@material-ui/core/Typography';
import axios from 'axios';
import Header from '../components/Header'
class ContainerPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = { fundModel: {} };
  }
  componentDidMount() {
    if (!window.location.pathname.includes("/fon/")) {
      return;
    }
    let urlArray = window.location.pathname.split('/');
    let fundCode = urlArray[urlArray.length - 1];
    axios.get(`http://localhost:5000/funds/` + fundCode)
      .then(res => {
        const fundModel = res.data;
        this.setState({ fundModel });
      })
  }

  render() {
    if (window.location.pathname.length <= 1) {
      return (
        <div>
          <Header />
        </div>)
    }
    else {
      return (
        <div>
          <Header />
          <FundPage fundModel={this.state.fundModel} />
          <Box pt={1}>
            <Typography variant="body2" color="textSecondary" align="center">
              {'Copyright © Fon Fiyatları'}
            </Typography>
          </Box>
        </div>
      )
    }
  }
}

export default ContainerPage;