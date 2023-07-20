"use client";
import TopPage from "@/components/TopPage/TopPage";
import { Mail, MessageSquare, User } from "lucide-react";
import React from "react";

function UsersPage() {
  return (
    <main>
      <TopPage
        Icon={User}
        IconColor="text-purple-500 "
        Titlei18Label="Commons.User"
        BgIconColor="bg-purple-200/50"
        Descriptioni18Label="UsersPage.TopDescription"
      />
    </main>
  );
}

export default UsersPage;
