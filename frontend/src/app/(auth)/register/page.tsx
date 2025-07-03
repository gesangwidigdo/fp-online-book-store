"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import Link from "next/link";
import { useRouter } from "next/navigation";

// 1. Define the schema for registration form
const formSchema = z.object({
  name: z.string().min(1, "Name is required"),
  email: z
    .string()
    .email("Please enter a valid email address")
    .min(1, "Email is required"),
  password: z.string().min(6, "Password must be at least 6 characters"),
  address: z.string().min(1, "Address is required"),
  gender: z.enum(["Male", "Female"], {
    errorMap: () => ({ message: "Gender is required" }),
  }),
  phoneNumber: z
    .string()
    .min(10, "Phone number must be at least 10 characters")
    .max(15, "Phone number cannot exceed 15 characters"),
  birthDate: z.string().min(1, "Birthdate is required"),
});

export default function RegisterPage() {
  const router = useRouter();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: "",
      email: "",
      password: "",
      address: "",
      gender: "Male", // Default value for gender
      phoneNumber: "",
      birthDate: "",
    },
  });

  // 2. Define the submit handler.
  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    try {
      console.log(values);
      const response = await fetch("http://localhost:5000/api/user/register", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          name: values.name,
          email: values.email,
          password: values.password,
          address: values.address,
          gender: values.gender,
          phone_number: values.phoneNumber,
          birth_date: values.birthDate,
        }),
      });

      if (!response.ok) {
        throw new Error("Failed to register");
      }

      const data = await response.json();
      console.log(data);
      router.push("/login");
    } catch (error) {
      console.error("Error during registration:", error);
    }
  };

  return (
    <div className="max-w-sm mx-auto mt-7">
      <h1 className="font-extrabold text-5xl mb-5">Register</h1>

      <Form {...form}>
        <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-2">
          {/* Name Field */}
          <FormField
            control={form.control}
            name="name"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Name</FormLabel>
                <FormControl>
                  <Input placeholder="Enter your name" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Email Field */}
          <FormField
            control={form.control}
            name="email"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Email</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Enter your email"
                    type="email"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Password Field */}
          <FormField
            control={form.control}
            name="password"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Password</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Enter your password"
                    type="password"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Address Field */}
          <FormField
            control={form.control}
            name="address"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Address</FormLabel>
                <FormControl>
                  <Input placeholder="Enter your address" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Gender Field */}
          <FormField
            control={form.control}
            name="gender"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Gender</FormLabel>
                <FormControl>
                  <select
                    className="form-select block w-full mt-1 border rounded-lg"
                    {...field}
                  >
                    <option value="Male">Male</option>
                    <option value="Female">Female</option>
                  </select>
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Phone Number Field */}
          <FormField
            control={form.control}
            name="phoneNumber"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Phone Number</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Enter your phone number"
                    type="tel"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Birth Date Field */}
          <FormField
            control={form.control}
            name="birthDate"
            render={({ field }) => (
              <FormItem>
                <FormLabel>Birth Date</FormLabel>
                <FormControl>
                  <Input
                    placeholder="Enter your birth date"
                    type="date"
                    {...field}
                  />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />

          {/* Submit Button */}
          <Button type="submit">Register</Button>
        </form>
      </Form>
      <p className="mt-5">
        Sudah memiliki akun?{" "}
        <Link href="/login">
          <span className="text-blue-500 underline cursor-pointer">
            Login sekarang!
          </span>
        </Link>
      </p>
    </div>
  );
}

