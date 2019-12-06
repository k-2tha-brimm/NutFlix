import React from 'react';
import './App.css';

export default class App extends React.Component {f

  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: ''
    }
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  update(field) {
    return e => this.setState({
      [field]: e.currentTarget.value
    });
  }

  async handleSubmit(e) {
    e.preventDefault();
    const user = Object.assign({}, this.state);
    return await fetch('/login', {
      method: 'POST',
      body: JSON.stringify(user),
      headers: {
        'Content-Type': 'application/json'
      }
    }).then(res => { return res.json() })
  }

  render() {
    return (
      <div className="App">
        <form onSubmit={this.handleSubmit}>
          <label /> Email
          <input type="text"
                 value={this.state.email}
                 onChange={this.update('email')}/>
  
          <label /> Password
          <input type="password"
                 value={this.state.password}
                 onChange={this.update('password')}/>
        <input className="session-submit" type="submit" value={this.props.formType} />
        </form>
      </div>
    );
  }
}
