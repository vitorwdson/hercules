import "../styles/global.css";
import Providers from "../utils/providers";

export const metadata = {
  title: "Bug Tracker",
  description: "Track issues and manage projects.",
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
