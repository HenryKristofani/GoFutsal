"use client"

import { Button } from "@/components/ui/button"
import { LogIn, UserPlus, Menu, LogOut, User } from "lucide-react"
import { useState, useEffect } from "react"
import { AuthDialog } from "./auth-dialog"
import { DropdownMenu, DropdownMenuContent, DropdownMenuItem, DropdownMenuTrigger, DropdownMenuSeparator } from "@/components/ui/dropdown-menu"
import { AuthAPI, type User as UserType } from "@/lib/auth"

export function Header() {
  const [showAuthDialog, setShowAuthDialog] = useState(false)
  const [authMode, setAuthMode] = useState<"login" | "register">("login")
  const [mounted, setMounted] = useState(false)
  const [user, setUser] = useState<UserType | null>(null)
  const [isAuthenticated, setIsAuthenticated] = useState(false)

  useEffect(() => {
    setMounted(true)
    // Check authentication status
    const authUser = AuthAPI.getCurrentUser()
    const authenticated = AuthAPI.isAuthenticated()
    setUser(authUser)
    setIsAuthenticated(authenticated)
  }, [])

  const handleAuthClick = (mode: "login" | "register") => {
    setAuthMode(mode)
    setShowAuthDialog(true)
  }

  const handleLogout = () => {
    AuthAPI.logout()
    setUser(null)
    setIsAuthenticated(false)
  }

  return (
    <>
      <header className="border-b border-border bg-card/50 backdrop-blur-sm sticky top-0 z-50">
        <div className="container mx-auto px-3 sm:px-4 lg:px-6 py-3 sm:py-4 max-w-7xl">
          <div className="flex items-center justify-between">
            <div className="flex items-center gap-2 sm:gap-3">
              <div className="w-8 h-8 sm:w-10 sm:h-10 rounded-lg bg-primary flex items-center justify-center">
                <svg
                  viewBox="0 0 24 24"
                  fill="none"
                  className="w-4 h-4 sm:w-6 sm:h-6 text-primary-foreground"
                  stroke="currentColor"
                  strokeWidth="2"
                >
                  <circle cx="12" cy="12" r="9" />
                  <path d="M12 3v18M3 12h18" />
                </svg>
              </div>
              <h1 className="text-lg sm:text-xl lg:text-2xl font-bold">
                Futsal<span className="text-primary">Arena</span>
              </h1>
            </div>

            <nav className="hidden lg:flex items-center gap-4 xl:gap-6">
              <a href="#lapangan" className="text-sm xl:text-base font-medium hover:text-primary transition-colors">
                Lapangan
              </a>
              <a href="#ketersediaan" className="text-sm xl:text-base font-medium hover:text-primary transition-colors">
                Ketersediaan
              </a>
              <a href="#tentang" className="text-sm xl:text-base font-medium hover:text-primary transition-colors">
                Tentang
              </a>
            </nav>

            <div className="flex items-center gap-1 sm:gap-2">
              <div className="hidden md:flex items-center gap-1 lg:gap-2">
                {isAuthenticated && user ? (
                  <>
                    <div className="flex items-center gap-2 px-2 py-1 rounded-md bg-muted/50">
                      <User className="w-3 h-3 sm:w-4 sm:h-4" />
                      <span className="text-xs sm:text-sm font-medium">
                        {user.username}
                      </span>
                      <span className="text-xs text-muted-foreground bg-primary/10 px-1.5 py-0.5 rounded">
                        {user.role}
                      </span>
                    </div>
                    <Button variant="ghost" size="sm" onClick={handleLogout} className="text-xs sm:text-sm">
                      <LogOut className="w-3 h-3 sm:w-4 sm:h-4 mr-1 sm:mr-2" />
                      <span className="hidden sm:inline">Keluar</span>
                      <span className="sm:hidden">Out</span>
                    </Button>
                  </>
                ) : (
                  <>
                    <Button variant="ghost" size="sm" onClick={() => handleAuthClick("login")} className="text-xs sm:text-sm">
                      <LogIn className="w-3 h-3 sm:w-4 sm:h-4 mr-1 sm:mr-2" />
                      <span className="hidden sm:inline">Masuk</span>
                      <span className="sm:hidden">Login</span>
                    </Button>
                    <Button size="sm" onClick={() => handleAuthClick("register")} className="text-xs sm:text-sm">
                      <UserPlus className="w-3 h-3 sm:w-4 sm:h-4 mr-1 sm:mr-2" />
                      <span className="hidden sm:inline">Daftar</span>
                      <span className="sm:hidden">Sign</span>
                    </Button>
                  </>
                )}
              </div>

              {mounted && (
                <DropdownMenu>
                  <DropdownMenuTrigger asChild className="md:hidden" suppressHydrationWarning>
                    <Button variant="ghost" size="icon" className="w-8 h-8 sm:w-10 sm:h-10">
                      <Menu className="w-4 h-4 sm:w-5 sm:h-5" />
                    </Button>
                  </DropdownMenuTrigger>
                  <DropdownMenuContent align="end" className="w-44 sm:w-48">
                    <DropdownMenuItem className="lg:hidden">
                      <a href="#lapangan" className="flex items-center w-full">
                        Lapangan
                      </a>
                    </DropdownMenuItem>
                    <DropdownMenuItem className="lg:hidden">
                      <a href="#ketersediaan" className="flex items-center w-full">
                        Ketersediaan
                      </a>
                    </DropdownMenuItem>
                    <DropdownMenuItem className="lg:hidden">
                      <a href="#tentang" className="flex items-center w-full">
                        Tentang
                      </a>
                    </DropdownMenuItem>
                    
                    <DropdownMenuSeparator className="lg:hidden" />
                    
                    {isAuthenticated && user ? (
                      <>
                        <DropdownMenuItem disabled>
                          <User className="w-4 h-4 mr-2" />
                          {user.username} ({user.role})
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={handleLogout}>
                          <LogOut className="w-4 h-4 mr-2" />
                          Keluar
                        </DropdownMenuItem>
                      </>
                    ) : (
                      <>
                        <DropdownMenuItem onClick={() => handleAuthClick("login")}>
                          <LogIn className="w-4 h-4 mr-2" />
                          Masuk
                        </DropdownMenuItem>
                        <DropdownMenuItem onClick={() => handleAuthClick("register")}>
                          <UserPlus className="w-4 h-4 mr-2" />
                          Daftar
                        </DropdownMenuItem>
                      </>
                    )}
                  </DropdownMenuContent>
                </DropdownMenu>
              )}
            </div>
          </div>
        </div>
      </header>

      <AuthDialog open={showAuthDialog} onOpenChange={setShowAuthDialog} mode={authMode} onModeChange={setAuthMode} />
    </>
  )
}
