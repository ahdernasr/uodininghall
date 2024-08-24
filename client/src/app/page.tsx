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
import { useToast } from "@/components/ui/use-toast";

// const fetcher = (url: string) =>
//   fetch(`http://localhost:4000/${url}`).then((res) => {
//     if (!res.ok) {
//       throw new Error("An error occurred while fetching the data.");
//     }
//     return res.text();
//   });

// const URL = "http://127.0.0.1:4000"
const URL = "https://uodininghall.live/uodininghall-server";

export default function Home() {
  const { toast } = useToast();
  const [email, setEmail] = useState<string>("");
  const [lock, setLock] = useState(false);

  const handleSubmit = async (email: string) => {
    if (lock) {
      toast({
        title: "Subscribe cooldown",
        description: "Please wait a moment before clicking again.",
        duration: 3000,
      });
      return;
    }

    try {
      if (!validateEmail(email)) {
        toast({
          title: "Error: Invalid email address",
          description:
            "Please make sure you email address is valid and try again.",
          duration: 3000,
        });
        return;
      }

      const response = await fetch(`${URL}/api/subscribe`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email: email }),
      });
      if (!response.ok) {
        console.log(response);
        throw new Error("Network response was not ok");
      }
      const data = await response.text();

      toast({
        title: "Subscription successful!",
        description:
          "You will now recieve the UO Dining Hall menu everyday at 6am.",
        duration: 3000,
      });

      setEmail("");
      lockButton(10);

    } catch (error) {
      toast({
        title: "Error: Could not subscribe",
        description: "Please try again.",
        duration: 3000,
      });

      console.log(error)

      setEmail("");
      lockButton(3)
    }
  };

  const validateEmail = (email: string): boolean => {
    // Regular expression for basic email validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    // Check if the email matches the regex pattern
    if (!emailRegex.test(email)) {
      return false;
    }

    // Check if the email length is within the acceptable range
    if (email.length > 254) {
      return false;
    }

    // Check if the local part (before @) and domain part (after @) are within acceptable lengths
    const [localPart, domainPart] = email.split("@");
    if (localPart.length > 64 || domainPart.length > 255) {
      return false;
    }

    // Ensure the domain part contains at least one dot and the top-level domain is valid
    const domainParts = domainPart.split(".");
    if (
      domainParts.length < 2 ||
      domainParts.some((part) => part.length === 0)
    ) {
      return false;
    }

    // Basic check to see if the top-level domain has at least two characters
    const topLevelDomain = domainParts[domainParts.length - 1];
    if (topLevelDomain.length < 2) {
      return false;
    }

    return true;
  };

  const lockButton = (duration: number) => {
    setLock(true);
    setTimeout((): void => {
      setLock(false);
    }, duration * 1000);
  };

  return (
    <>
      <main className="flex flex-col justify-center items-center gap-y-5">
        <h1 className="text-[4rem] text-[#8F001A] font-bold ">
          UO Dining Hall
        </h1>
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
                  Subscribe to recieve uOttawa&apos;s Dining Hall menu everyday
                  by email at 6:00am.
                </CardDescription>
              </CardHeader>
              <CardContent className="space-y-2">
                <div className="space-y-1">
                  <Label htmlFor="email">Email</Label>
                  <Input
                    value={email}
                    id="email"
                    placeholder="user@uottawa.ca"
                    onChange={(e) => setEmail(e.target.value)}
                  />
                </div>
              </CardContent>
              <CardFooter>
                <Button
                  onClick={() => handleSubmit(email)}
                  className="bg-[#8F001A]"
                >
                  Subscribe
                </Button>
              </CardFooter>
            </Card>
          </TabsContent>
          <TabsContent value="password">
            <Card>
              <CardHeader>
                <CardTitle>Subscribe wth SMS</CardTitle>
                <CardDescription>
                  Subscribe to recieve uOttawa&apos;s Dining Hall menu everyday
                  by SMS at 6:00am.
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
          </Link>
          .
        </div>
      </main>
    </>
  );
}
