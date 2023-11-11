import { getClientSideLocale, getClientSideTranslation } from "@/i18n";
import "./globals.css";
import { NextIntlClientProvider } from "next-intl";
import { Frank_Ruhl_Libre } from "next/font/google";
import {} from "next/font/local";
import { cn } from "@/lib/utils";
import { Toaster } from "react-hot-toast";
const inter = Frank_Ruhl_Libre({
  subsets: ["latin"],
});

export default async function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <NextIntlClientProvider
        locale={getClientSideLocale()}
        messages={await getClientSideTranslation()}
      >
        <body className={cn(inter.className, "bg-white", "h-screen")}>
          <Toaster position="top-center" />
          {children}
        </body>
      </NextIntlClientProvider>
    </html>
  );
}
