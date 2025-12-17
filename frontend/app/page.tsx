"use client"

import { Header } from "@/components/header"
import { CourtList } from "@/components/court-list"
import { AvailabilityCalendar } from "@/components/availability-calendar"

export default function HomePage() {
  return (
    <div className="min-h-screen bg-background">
      <Header />
      <main className="container mx-auto px-3 sm:px-4 lg:px-6 py-6 sm:py-8 lg:py-12 max-w-7xl">
        <section className="mb-8 sm:mb-12 lg:mb-16">
          <div className="text-center mb-6 sm:mb-8 lg:mb-10">
            <h1 className="text-2xl sm:text-3xl md:text-4xl lg:text-5xl xl:text-6xl font-bold mb-3 sm:mb-4 lg:mb-6 text-balance leading-tight">
              Booking Lapangan Futsal <span className="text-primary">Online</span>
            </h1>
            <p className="text-muted-foreground text-sm sm:text-base lg:text-lg xl:text-xl max-w-xs sm:max-w-md lg:max-w-2xl xl:max-w-3xl mx-auto text-pretty px-2 sm:px-0">
              Sistem booking lapangan futsal yang mudah, cepat, dan terpercaya. Pilih lapangan favoritmu dan booking
              sekarang!
            </p>
          </div>
        </section>

        <CourtList />

        <section className="mt-12 sm:mt-16 lg:mt-20">
          <AvailabilityCalendar />
        </section>
      </main>
    </div>
  )
}
