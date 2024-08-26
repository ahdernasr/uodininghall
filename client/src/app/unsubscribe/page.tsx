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
import { cn } from "@/lib/utils";

import { useState } from "react";
import { useToast } from "@/components/ui/use-toast";

// const URL = "http://127.0.0.1:4000"
const URL = "https://uodininghall.live/uodininghall-server";

export default function Unsubscribe() {
  const { toast } = useToast();
  const [email, setEmail] = useState<string>("");

  const handleSubmit = async (email: string) => {
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

      const response = await fetch(`${URL}/api/unsubscribe`, {
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

      toast({
        title: "Subscription removed.",
        description: "You will now stop recieving emails from us.",
        duration: 3000,
      });

      setEmail("");
    } catch (error) {
      toast({
        title: "Error: Could not unsubscribe",
        description: "Please try again.",
        duration: 3000,
      });

      setEmail("");
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

  return (
    <main className="flex flex-col justify-center items-center gap-y-5">
      <h1 className="text-[2.75rem] sm:text-[4rem] text-[#8F001A] font-bold ">UO Dining Hall</h1>
      <Card className="mt-[3%]">
        <CardHeader>
          <CardTitle>Unsubscribe</CardTitle>
          <CardDescription>
            Stop recieving the UO Dining Hall menu daily.
          </CardDescription>
        </CardHeader>
        <CardContent className="space-y-2">
          <div className="space-y-1">
            <Label htmlFor="email">Email</Label>
            <Input
              id="email"
              placeholder="user@uottawa.ca"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
        </CardContent>
        <CardFooter>
          <Button onClick={() => handleSubmit(email)} className="bg-[#8F001A]">
            Confirm
          </Button>
        </CardFooter>
      </Card>
      <div className={`m-2 ${cn("text-[0.8rem] text-muted-foreground")}`}>
        Want to subscribe? Click{" "}
        <Link href="/" className="underline">
          here
        </Link>
        .
      </div>
    </main>
  );
}
