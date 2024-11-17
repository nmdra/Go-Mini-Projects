## Generating Self-Signed SSL Certificates

To generate SSL certificates for local development and bypass browser warnings:

1. **Run the following command** to generate a self-signed SSL certificate and private key:

   ```bash
   openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /ssl/self-signed.key -out /ssl/self-signed.crt
   ```

2. **Move the generated files** (`self-signed.crt` and `self-signed.key`) to the `/ssl` folder in your project directory.

3. **Use these files** in your server configuration (e.g., Nginx) to enable HTTPS on your local development server.

4. **Bypass Browser Warning**:
   - **Chrome**: Click on "Advanced" on the SSL warning page and then click "Proceed to localhost (unsafe)".
   - **Firefox**: Click "Advanced" and then "Accept the Risk and Continue".
   - **Safari**: Go to **Preferences > Privacy** and add your domain to **Certificates** to trust it.

   Alternatively, for development purposes, you can import the self-signed certificate into your browser to avoid the warning on every visit.