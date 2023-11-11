import { NextRequest } from "next/server";

export function middleware(request: NextRequest) {

}

export const config = {
// do not localize next.js paths
  matcher: ["/((?!api|_next/static|_next/image|assets|favicon.ico|sw.js).*)",],
};