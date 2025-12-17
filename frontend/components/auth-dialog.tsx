"use client"

import type React from "react"

import { useState } from "react"
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
  const [email, setEmail] = useState("")
  const [password, setPassword] = useState("")
  const [name, setName] = useState("")
  const [phone, setPhone] = useState("")

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()
    // TODO: Implement authentication logic
    console.log("Auth submitted:", { email, password, name, phone, mode })
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
          {mode === "register" && (
            <>
              <div className="space-y-1.5 sm:space-y-2">
                <Label htmlFor="name" className="text-sm sm:text-base">Nama Lengkap</Label>
                <Input
                  id="name"
                  type="text"
                  placeholder="Masukkan nama lengkap"
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  required
                  className="h-10 sm:h-11 text-sm sm:text-base"
                />
              </div>
              <div className="space-y-1.5 sm:space-y-2">
                <Label htmlFor="phone" className="text-sm sm:text-base">Nomor Telepon</Label>
                <Input
                  id="phone"
                  type="tel"
                  placeholder="08xxxxxxxxxx"
                  value={phone}
                  onChange={(e) => setPhone(e.target.value)}
                  required
                  className="h-10 sm:h-11 text-sm sm:text-base"
                />
              </div>
            </>
          )}

          <div className="space-y-1.5 sm:space-y-2">
            <Label htmlFor="email" className="text-sm sm:text-base">Email</Label>
            <Input
              id="email"
              type="email"
              placeholder="nama@email.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
              className="h-10 sm:h-11 text-sm sm:text-base"
            />
          </div>

          <div className="space-y-1.5 sm:space-y-2">
            <Label htmlFor="password" className="text-sm sm:text-base">Password</Label>
            <Input
              id="password"
              type="password"
              placeholder="••••••••"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
              className="h-10 sm:h-11 text-sm sm:text-base"
            />
          </div>

          <Button type="submit" className="w-full h-10 sm:h-11 text-sm sm:text-base" size="lg">
            {mode === "login" ? (
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
