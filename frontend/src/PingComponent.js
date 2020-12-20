import React, { Component } from 'react'
import axios from 'axios';
class PingComponent extends Component {

    constructor() {
        super();
        this.state = {
            pong: 'pending'
        }
    }

    componentWillMount() {
        axios.get('http://localhost:8080/ping')
            .then((response) => {
                this.setState(() => {
                    return { pong: response.data }
                })
            })
            .catch(function(error) {
                console.log(error);
            });
    }

    render() {
        return <h1>Ping { this.state.pong }</h1>;
    }
}

export default PingComponent;