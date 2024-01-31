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
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

import { cn } from "@/lib/utils";
import { useState } from "react";

export const ENDPOINT = "http://localhost:4000";

const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((res) => {
    if (!res.ok) {
      throw new Error("An error occurred while fetching the data.");
    }
    return res.text();
  });

export default function Home() {
  const [ email, setEmail ] = useState<string>("")
  // const [ subscriptions, setSubscriptions ] = useState<string>('')

  // Registers every website visit
  const { } = useSWR("api/track-visit", fetcher);

  const handleSubmit = async (email: string) => {
    try {
      const response = await fetch(`${ENDPOINT}/api/subscribe`, {
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
      {/* <div
        className={`flex justify-center items-center flex-col m-2 ${cn(
          "text-[0.8rem] text-muted-foreground"
        )}`}
      >
        <p>Subscriptions: {subscriptions}</p>
      </div> */}
      <Tabs defaultValue="account" className="w-[400px]">
        <TabsList className="grid w-full grid-cols-2">
          <TabsTrigger value="account">Email</TabsTrigger>
          <TabsTrigger value="password">SMS</TabsTrigger>
        </TabsList>
        <TabsContent value="account">
          <Card>
            <CardHeader>
              <CardTitle>Subscribe with Email</CardTitle>
              <CardDescription>
                Subscribe to recieve uOttawa's Dining Hall menu everyday by
                email at 6:00am.
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-2">
              <div className="space-y-1">
                <Label htmlFor="email">Email</Label>
                <Input id="email" placeholder="user@uottawa.ca" onChange={(e) => setEmail(e.target.value)}/>
              </div>
            </CardContent>
            <CardFooter>
              <Button onClick={() => handleSubmit(email)}>Subscribe with email</Button>
            </CardFooter>
          </Card>
        </TabsContent>
        <TabsContent value="password">
          <Card>
            <CardHeader>
              <CardTitle>Subscribe wth SMS</CardTitle>
              <CardDescription>
                Subscribe to recieve uOttawa's Dining Hall menu everyday by SMS
                at 6:00am.
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-2">
              <div className="space-y-1">
                <Label htmlFor="phone">Phone Number</Label>
                <Input id="phone" type="phone" />
              </div>
            </CardContent>
            <CardFooter>
              <Button variant="secondary" disabled={true}>
                Coming soon...
              </Button>
            </CardFooter>
          </Card>
        </TabsContent>
      </Tabs>
      <div className={`m-2 ${cn("text-[0.8rem] text-muted-foreground")}`}>
        Want to unsubscribe? Click{" "}
        <Link href="/unsubscribe" className="underline">
          here
        </Link>.
      </div>
    </main>
  );
}
