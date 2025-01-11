This is a [Next.js](https://nextjs.org) project bootstrapped with [`create-next-app`](https://nextjs.org/docs/app/api-reference/cli/create-next-app).

## Getting Started

First, run the development server:

```bash
npm run dev
# or
yarn dev
# or
pnpm dev
# or
bun dev
```

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

You can start editing the page by modifying `app/page.tsx`. The page auto-updates as you edit the file.

This project uses [`next/font`](https://nextjs.org/docs/app/building-your-application/optimizing/fonts) to automatically optimize and load [Geist](https://vercel.com/font), a new font family for Vercel.

## Learn More

To learn more about Next.js, take a look at the following resources:

- [Next.js Documentation](https://nextjs.org/docs) - learn about Next.js features and API.
- [Learn Next.js](https://nextjs.org/learn) - an interactive Next.js tutorial.

You can check out [the Next.js GitHub repository](https://github.com/vercel/next.js) - your feedback and contributions are welcome!

## Deploy on Vercel

The easiest way to deploy your Next.js app is to use the [Vercel Platform](https://vercel.com/new?utm_medium=default-template&filter=next.js&utm_source=create-next-app&utm_campaign=create-next-app-readme) from the creators of Next.js.

Check out our [Next.js deployment documentation](https://nextjs.org/docs/app/building-your-application/deploying) for more details.

# Google Cloud Project and OAuth 2.0 Client ID Setup Guide

This guide walks you through the steps to create a Google Cloud project and generate an OAuth 2.0 Client ID for authenticating users in your application.

## Prerequisites

- A Google account.
- Access to Google Cloud Console.
- Internet connection.

## Steps to Create a Google Cloud Project

1. **Go to the Google Cloud Console:**
   Open the [Google Cloud Console](https://console.cloud.google.com/projectcreate?inv=1&invt=AbmdTQ) in your web browser.

2. **Create a New Project:**
    - Click on **"Create Project"**.
    - Choose a name for your project and select a billing account (if necessary).
    - Click on **"Create"** to set up the project.

3. **Select Your Project:**
   After your project is created, you will be redirected to the **Project Dashboard**. Make sure that your new project is selected in the top dropdown of the Google Cloud Console.

## Steps to Create an OAuth 2.0 Client ID

1. **Go to the OAuth 2.0 Credentials Page:**
   Open the [Google API Credentials page](https://console.cloud.google.com/apis/credentials?inv=1&invt=AbmdTA&project=elaborate-truth-447408-r9) in your web browser.

2. **Create OAuth Client ID:**
    - Click **"Create Credentials"** and select **"OAuth 2.0 Client ID"**.
    - If you haven't configured the OAuth consent screen, you will be prompted to do so. You must provide information such as the **application name** and **support email**.
    - After configuring the consent screen, proceed with creating the client ID.

3. **Configure OAuth Client ID:**
    - Select the application type (for web, choose **Web Application**).
    - Add the **authorized JavaScript origins** (for example: `http://localhost:8080`).
    - Add **authorized redirect URIs** where the OAuth server will send responses (for example: `http://localhost:8080/callback`).

4. **Generate the OAuth 2.0 Credentials:**
    - Click on **"Create"** to generate the Client ID.
    - Save the **Client ID** and **Client Secret** for use in your application.

## Integrating OAuth 2.0 with Your Application
