"use client";

import React from 'react';
import { GoogleLogin, CredentialResponse, GoogleOAuthProvider } from '@react-oauth/google';

const GoogleLoginButton: React.FC = () => {
    const GOOGLE_CLIENT_ID = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '';

    const onSuccess = (res: CredentialResponse) => {
        if (res.credential) {
            console.log('Login Success:', res.credential);
        } else {
            console.error('No credentials received');
        }
    };

    const handleLoginFailure = (error?: unknown) => {
        if (error instanceof DOMException && error.name === 'AbortError') {
            console.error('Login aborted by user or browser environment issue.');
        } else {
            console.error('Unexpected error occurred:', error);
        }
    };


    return (
        <GoogleOAuthProvider clientId={GOOGLE_CLIENT_ID}>
            <div>
                <h3>Login with Google</h3>
                <GoogleLogin
                    onSuccess={onSuccess}
                    onError={handleLoginFailure}
                    useOneTap={false}
                    logo_alignment='center'
                />
            </div>
        </GoogleOAuthProvider>
    );
};

export default GoogleLoginButton;
