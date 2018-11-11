class App extends React.Component {
  render() {
      return (<LoggedIn />);
  }
}

class Home extends React.Component {
  render() {
    return (
      <div className="container">
        <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
          <h1>Magic</h1>
          <p>Super Cool Stuff</p>
          <p>Sign in to get access </p>
          <a onClick={this.authenticate} className="btn btn-primary btn-lg btn-login btn-block">Sign In</a>
        </div>
      </div>
    )
  }
}

class LoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      starttime: [],
      uptime: [],
      clientip: [],
      uniqueip: []
    };
    
    this.serverStartTimeRequest = this.serverStartTimeRequest.bind(this);
    this.serverUptimeRequest = this.serverUptimeRequest.bind(this);
    this.serverClientIPRequest = this.serverClientIPRequest.bind(this);
    this.serverUniqueIPRequest = this.serverUniqueIPRequest.bind(this);

  }

  serverStartTimeRequest() {
    $.get("http://a66d6a982dca511e89b1f12e4c7911df-1486875010.us-east-1.elb.amazonaws.com/api/starttime", res => {
      this.setState({
        starttime: res
      });
    });
  }
  
  serverUptimeRequest() {
    $.get("http://a66d6a982dca511e89b1f12e4c7911df-1486875010.us-east-1.elb.amazonaws.com/api/uptime", res => {
      this.setState({
        uptime: res
      });
    });
  }

  serverClientIPRequest() {
    $.get("http://a66d6a982dca511e89b1f12e4c7911df-1486875010.us-east-1.elb.amazonaws.com/api/ip", res => {
      this.setState({
        clientip: res
      });
    });
  }

  serverUniqueIPRequest() {
    $.get("http://a66d6a982dca511e89b1f12e4c7911df-1486875010.us-east-1.elb.amazonaws.com/api/uniqueip", res => {
      this.setState({
        uniqueip: res
      });
    });
  }

  componentDidMount() {
    this.serverStartTimeRequest();
    this.serverUptimeRequest();
    this.serverClientIPRequest();
    this.serverUniqueIPRequest();
  }

  render() {
    return (
      <div className="container">
        <div className="col-lg-12">
          <br />
          <span className="pull-right"><a onClick={this.logout}>Log out</a></span>
          <h2>Something Completely Different</h2>
          <p>The Droids You Are Looking For:</p>
          <div className="row">
            <div className="container">
              <Starttime message={this.state.starttime.message} starttime={this.state.starttime.starttime}/>
            </div>
          </div>
          <div className="row">
            <div className="container">
              <Uptime message={this.state.uptime.message} uptime={this.state.uptime.uptime}/>
            </div>
          </div>
          <div className="row">
            <div className="container">
              <ClientIP message={this.state.clientip.message} clientip={this.state.clientip.clientip}/>
            </div>
          </div>
          <div className="row">
            <div className="container">
              <Uniqueip message={this.state.uniqueip.message} uniqueip={this.state.uniqueip.uniqueip}/>
            </div>
          </div>
        </div>
      </div>
    )
  }
}

class Starttime extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.message}</div>
          <div className="panel-body">{this.props.starttime}</div>
        </div>
      </div>
    )
  }
}

class Uptime extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.message}</div>
          <div className="panel-body">{this.props.uptime}</div>
        </div>
      </div>
    )
  }
}

class ClientIP extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.message}</div>
          <div className="panel-body">{this.props.clientip}</div>
        </div>
      </div>
    )
  }
}

class Uniqueip extends React.Component {
  constructor(props) {
    super(props);
  }

  render() {
    return (
      <div className="col-xs-4">
        <div className="panel panel-default">
          <div className="panel-heading">{this.props.message}</div>
          <div className="panel-body">{this.props.uniqueip}</div>
        </div>
      </div>
    )
  }
}

ReactDOM.render(<App />, document.getElementById('app'));
