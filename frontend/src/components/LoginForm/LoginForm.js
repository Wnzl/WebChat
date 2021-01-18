import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { object, string, ref } from 'yup';
import { Card, CardContent, CardHeader, Typography, withStyles } from '@material-ui/core';
import Login from "../Login/Login";

const styles = theme => ({
    container: {
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    },
    card: {
        width: '500px',
        margin: '20px',
    },
    headerContainer: {
        display: 'flex',
        alignItems: 'center',
    },
    cardHeader: {
        marginLeft: '10px',
        color: theme.palette.primary.main,
        cursor: 'pointer',
        transition: 'all .2s',
        '&:hover': {
            color: '#585858',
        },
    },
    cardEnabled: {
        color: theme.palette.primary.main,
        fontWeight: 'bold',
    },
});

class LoginForm extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isSignInOpen: true,
        };
        this.renderForm = this.renderForm.bind(this);
        this.handleSignInOpenClick = this.handleSignInOpenClick.bind(this);
        this.handleSignUpOpenClick = this.handleSignUpOpenClick.bind(this);
    }

    handleSignInOpenClick() {
        this.setState({
            isSignInOpen: true,
        });
    }
    handleSignUpOpenClick() {
        this.setState({
            isSignInOpen: false,
        });
    }

    renderForm() {
        const { startLogin, startRegister } = this.props;
        const { isSignInOpen } = this.state;

        const validationSignInSchema = object().shape({
            email: string()
                .email('E-mail is not valid')
                .required('E-mail is required'),
            password: string()
                .min(6, 'Password must be longer or equal to 6 characters!')
                .max(66, 'Password has to be shorter than 66 characters')
                .required('Password is required'),
        });

        const validationSignUpSchema = object().shape({
            email: string()
                .email('E-mail is not valid')
                .required('E-mail is required'),
            password: string()
                .min(6, 'Password must be longer or equal to 6 characters!')
                .max(66, 'Password has to be shorter than 66 characters')
                .required('Password is required'),
            passwordConfirmation: string()
                .oneOf([ref('password'), null], "Passwords don't match")
                .required('Password confirm is required'),
        });
        if (isSignInOpen) {
            return <Login handleLogin={startLogin} validationSchema={validationSignInSchema} />;
        }
        return (
            <Login
                handleLogin={startRegister}
                validationSchema={validationSignUpSchema}
                status="signUp"
            />
        );
    }

    render() {
        const { classes } = this.props;
        const { isSignInOpen } = this.state;
        return (
            <div className={classes.container}>
                <Card className={classes.card}>
                    <CardHeader
                        title={
                            <div className={classes.headerContainer}>
                                <Typography
                                    variant="subtitle1"
                                    className={`${classes.cardHeader} ${
                                        isSignInOpen ? classes.cardEnabled : ''
                                    }`}
                                    onClick={this.handleSignInOpenClick}
                                >
                                    Login
                                </Typography>
                                <Typography
                                    variant="subtitle1"
                                    className={`${classes.cardHeader} ${
                                        isSignInOpen ? '' : classes.cardEnabled
                                    }`}
                                    onClick={this.handleSignUpOpenClick}
                                >
                                    SignUp
                                </Typography>
                            </div>
                        }
                    />
                    <CardContent>{this.renderForm()}</CardContent>
                </Card>
            </div>
        );
    }
}

LoginForm.propTypes = {
    startLogin: PropTypes.func.isRequired,
    startRegister: PropTypes.func.isRequired,
    classes: PropTypes.shape({}).isRequired,
};

export default withStyles(styles)(LoginForm);
