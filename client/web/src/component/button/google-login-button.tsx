import { GoogleLogin, GoogleOAuthProvider } from "@react-oauth/google";

export default function GoogleLoginButton() {
  const clientId = process.env.REACT_APP_GOOGLE_OAUTH_ID as string;

  return (
    <GoogleOAuthProvider clientId={clientId}>
      <GoogleLogin
        onSuccess={(credentialResponse) => {
          console.log({ credentialResponse });
        }}
        onError={() => {
          console.log("Google Login Failed");
        }}
      />
    </GoogleOAuthProvider>
  );
}
