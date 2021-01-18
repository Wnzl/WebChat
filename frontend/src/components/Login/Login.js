import React from "react";
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import {Form, Formik} from "formik";
import {FormHelperText, TextField} from "@material-ui/core";

const styles = () => ({
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
    cardHeader: {
        textAlign: 'center',
        color: '#585858',
    },
    actionsContainer: {
        marginTop: '20px',
        marginBottom: '20px',
    },
});

const initialValues = {
    email: '',
    password: '',
    passwordConfirmation: '',
};

const Login = ({ classes, validationSchema, handleLogin, status }) => (
    <Formik
        enableReinitialize
        initialValues={initialValues}
        validationSchema={validationSchema}
        onSubmit={(values, { setSubmitting }) => {
            setSubmitting(false);
            handleLogin(values);
        }}
        render={({ touched, errors, handleChange, handleSubmit, handleBlur, isValid, values }) => (
            <Form>
                <div>
                    <TextField
                        name="email"
                        value={values.email}
                        label="Email"
                        fullWidth
                        placeholder="Enter your email"
                        margin="normal"
                        onBlur={handleBlur}
                        onChange={handleChange}
                    />
                    {errors.email &&
                    touched.email && <FormHelperText error>{errors.email}</FormHelperText>}
                </div>
                <div>
                    <TextField
                        name="password"
                        value={values.password}
                        label="Password"
                        type="password"
                        placeholder="Enter your password"
                        fullWidth
                        onBlur={handleBlur}
                        margin="normal"
                        onChange={handleChange}
                    />
                    {errors.password &&
                    touched.password && (
                        <FormHelperText error>{errors.password}</FormHelperText>
                    )}
                </div>
                {status === 'signUp' && (
                    <div>
                        <TextField
                            name="passwordConfirmation"
                            value={values.passwordConfirmation}
                            label="Password confirm"
                            type="password"
                            placeholder="Confirm your password"
                            fullWidth
                            onBlur={handleBlur}
                            margin="normal"
                            onChange={handleChange}
                        />
                        {errors.passwordConfirmation &&
                        touched.passwordConfirmation && (
                            <FormHelperText error>{errors.passwordConfirmation}</FormHelperText>
                        )}
                    </div>
                )}
                <div className={classes.actionsContainer}>
                    <Button
                        variant="contained"
                        color="primary"
                        fullWidth
                        type="submit"
                        onClick={handleSubmit}
                        disabled={!isValid}
                    >
                        {status === 'signUp' ? 'Sign Up' : 'Log In'}
                    </Button>
                </div>
            </Form>
        )}
    />
);

Login.defaultProps = {
    status: '',
};

Login.propTypes = {
    classes: PropTypes.shape({}).isRequired,
    validationSchema: PropTypes.shape({}).isRequired,
    handleLogin: PropTypes.func.isRequired,
    status: PropTypes.string,
};

export default withStyles(styles)(Login);
