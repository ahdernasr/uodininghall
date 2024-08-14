import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import { Toaster } from "@/components/ui/toaster";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "UO Dining Hall",
  description: "",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={`${inter.className} relate w-full h-full m-0 p-0 `}>
        {children}
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 1440 320"
          className="absolute bottom-0 left-0"
        >
          <path
            fill="#8F001A"
            fill-opacity="1"
            d="M0,192L34.3,181.3C68.6,171,137,149,206,160C274.3,171,343,213,411,213.3C480,213,549,171,617,176C685.7,181,754,235,823,250.7C891.4,267,960,245,1029,245.3C1097.1,245,1166,267,1234,256C1302.9,245,1371,203,1406,181.3L1440,160L1440,320L1405.7,320C1371.4,320,1303,320,1234,320C1165.7,320,1097,320,1029,320C960,320,891,320,823,320C754.3,320,686,320,617,320C548.6,320,480,320,411,320C342.9,320,274,320,206,320C137.1,320,69,320,34,320L0,320Z"
          ></path>
        </svg>
        <Toaster />
      </body>
    </html>
  );
}
