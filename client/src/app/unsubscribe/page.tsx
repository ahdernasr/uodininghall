"use client";

import { Button } from "@/components/ui/button";
import Link from "next/link";
import useSWR from "swr";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

import { useState } from "react";

export default function Unsubscribe() {
  const [email, setEmail] = useState<string>("");

  const handleSubmit = async (email: string) => {
    try {
      const response = await fetch(`http://localhost:4000/api/unsubscribe`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: email }),
      });
      if (!response.ok) {
        throw new Error("Network response was not ok");
      }
      const data = await response.text();
      // setSubscriptions(data)
    } catch (error) {
      console.error("Error:", error);
    }
  };

  return (
    <main>
      <Card>
        <CardHeader>
          <CardTitle>Unsubscribe</CardTitle>
          <CardDescription>
            Stop recieving the UO dining hall menu daily.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-2">
          <div className="space-y-1">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              placeholder="user@uottawa.ca"
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
        </CardContent>
        <CardFooter>
          <Button onClick={() => handleSubmit(email)}>
            Unsubscribe
          </Button>
        </CardFooter>
      </Card>
    </main>
  );
}
