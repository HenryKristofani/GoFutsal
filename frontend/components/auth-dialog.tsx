"use client"

import type React from "react"

import { useState } from "react"
import { AuthAPI, type LoginRequest, type RegisterRequest } from "@/lib/auth"
import { Button } from "@/components/ui/button"
import { Dialog, DialogContent, DialogDescription, DialogHeader, DialogTitle } from "@/components/ui/dialog"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { LogIn, UserPlus } from "lucide-react"

interface AuthDialogProps {
  open: boolean
  onOpenChange: (open: boolean) => void
  mode: "login" | "register"
  onModeChange: (mode: "login" | "register") => void
}

export function AuthDialog({ open, onOpenChange, mode, onModeChange }: AuthDialogProps) {
  const [username, setUsername] = useState("")
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState("")
  const [success, setSuccess] = useState("")

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    setIsLoading(true)
    setError("")
    setSuccess("")

    try {
      if (mode === "register") {
        const registerData: RegisterRequest = {
          username,
          email,
          password,
        }

        const result = await AuthAPI.register(registerData)
        
        if (result.success) {
          setSuccess(result.message)
          // Switch to login mode after successful registration
          setTimeout(() => {
            onModeChange("login")
            setSuccess("")
            setPassword("") // Clear password for security
          }, 2000)
        } else {
          setError(result.message)
        }
      } else {
        const loginData: LoginRequest = {
          username,
          password,
        }

        const result = await AuthAPI.login(loginData)
        
        if (result.success) {
          setSuccess(result.message)
          // Close dialog and reload to update auth state
          setTimeout(() => {
            onOpenChange(false)
            window.location.reload()
          }, 1000)
        } else {
          setError(result.message)
        }
      }
    } catch (err) {
      setError("Terjadi kesalahan yang tidak terduga")
    } finally {
      setIsLoading(false)
    }
  }

  return (
    <Dialog open={open} onOpenChange={onOpenChange}>
      <DialogContent className="w-[95vw] max-w-sm sm:max-w-md lg:max-w-lg mx-auto">
        <DialogHeader className="space-y-2 sm:space-y-3">
          <DialogTitle className="text-xl sm:text-2xl leading-tight">
            {mode === "login" ? "Masuk ke Akun" : "Buat Akun Baru"}
          </DialogTitle>
          <DialogDescription className="text-sm sm:text-base">
            {mode === "login"
              ? "Masuk untuk melakukan booking lapangan futsal"
              : "Daftar untuk mulai booking lapangan futsal"}
          </DialogDescription>
        </DialogHeader>

        <form onSubmit={handleSubmit} className="space-y-3 sm:space-y-4 mt-3 sm:mt-4">
          {error && (
            <div className="bg-destructive/10 text-destructive text-sm p-3 rounded-md border border-destructive/20">
              {error}
            </div>
          )}
          
          {success && (
            <div className="bg-green-50 text-green-700 text-sm p-3 rounded-md border border-green-200">
              {success}
            </div>
          )}

          <div className="space-y-1.5 sm:space-y-2">
            <Label htmlFor="username" className="text-sm sm:text-base">Username</Label>
            <Input
              id="username"
              type="text"
              placeholder="Masukkan username"
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              required
              disabled={isLoading}
              className="h-10 sm:h-11 text-sm sm:text-base"
            />
          </div>

          {mode === "register" && (
            <div className="space-y-1.5 sm:space-y-2">
              <Label htmlFor="email" className="text-sm sm:text-base">Email</Label>
              <Input
                id="email"
                type="email"
                placeholder="nama@email.com"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                required
                disabled={isLoading}
                className="h-10 sm:h-11 text-sm sm:text-base"
              />
            </div>
          )}

          <div className="space-y-1.5 sm:space-y-2">
            <Label htmlFor="password" className="text-sm sm:text-base">Password</Label>
            <Input
              id="password"
              type="password"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              disabled={isLoading}
              className="h-10 sm:h-11 text-sm sm:text-base"
            />
          </div>

          <Button type="submit" className="w-full h-10 sm:h-11 text-sm sm:text-base" size="lg" disabled={isLoading}>
            {isLoading ? (
              <>
                <div className="w-3 h-3 sm:w-4 sm:h-4 mr-1.5 sm:mr-2 animate-spin border-2 border-current border-t-transparent rounded-full" />
                {mode === "login" ? "Masuk..." : "Mendaftar..."}
              </>
            ) : mode === "login" ? (
              <>
                <LogIn className="w-3 h-3 sm:w-4 sm:h-4 mr-1.5 sm:mr-2" />
                Masuk
              </>
            ) : (
              <>
                <UserPlus className="w-3 h-3 sm:w-4 sm:h-4 mr-1.5 sm:mr-2" />
                Daftar
              </>
            )}
          </Button>

          <div className="text-center text-xs sm:text-sm">
            {mode === "login" ? (
              <p className="text-muted-foreground">
                Belum punya akun?{" "}
                <button
                  type="button"
                  onClick={() => onModeChange("register")}
                  className="text-primary hover:underline font-medium touch-target"
                >
                  Daftar sekarang
                </button>
              </p>
            ) : (
              <p className="text-muted-foreground">
                Sudah punya akun?{" "}
                <button
                  type="button"
                  onClick={() => onModeChange("login")}
                  className="text-primary hover:underline font-medium touch-target"
                >
                  Masuk
                </button>
              </p>
            )}
          </div>
        </form>
      </DialogContent>
    </Dialog>
  )
}
