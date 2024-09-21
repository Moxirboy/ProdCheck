'use client'

import { useState, useEffect } from 'react'
import { motion, useScroll, useTransform } from 'framer-motion'
import { BrowserRouter as Router, Route, Routes, Link, useLocation } from 'react-router-dom'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { Select } from "@/components/ui/select"
import { BarChart, Globe, Shield, AlertTriangle, ChevronRight, Check, ArrowRight, Search, Filter, Bell, DollarSign, Mail, Phone, MapPin } from 'lucide-react'

// Layout component
const Layout = ({ children }) => {
  const location = useLocation()
  const isLandingPage = location.pathname === '/'

  return (
    <div className="flex flex-col min-h-screen bg-gradient-to-b from-blue-50 to-white">
      <motion.header 
        className="sticky top-0 z-50 bg-white bg-opacity-90 backdrop-blur-md shadow-sm"
        initial={{ y: -100 }}
        animate={{ y: 0 }}
        transition={{ duration: 0.5 }}
      >
        <div className="container mx-auto px-4 py-4 flex items-center justify-between">
          <Link to="/" className="flex items-center space-x-2">
            <BarChart className="h-8 w-8 text-blue-600" />
            <span className="text-2xl font-bold text-gray-900">ProdCheck</span>
          </Link>
          <nav className="hidden md:flex space-x-8">
            <Link to="/platform" className="text-gray-600 hover:text-blue-600 transition-colors">Platform</Link>
            {isLandingPage && (
              <>
                <a href="#features" className="text-gray-600 hover:text-blue-600 transition-colors">Features</a>
                <a href="#pricing" className="text-gray-600 hover:text-blue-600 transition-colors">Pricing</a>
                <a href="#contact" className="text-gray-600 hover:text-blue-600 transition-colors">Contact</a>
              </>
            )}
            {!isLandingPage && (
              <>
                <Link to="/analyze" className="text-gray-600 hover:text-blue-600 transition-colors">Analyze Market</Link>
                <Link to="/dashboard" className="text-gray-600 hover:text-blue-600 transition-colors">Dashboard</Link>
              </>
            )}
          </nav>
          <div className="flex space-x-4">
            <Button variant="outline">Sign In</Button>
            <Button>Sign Up</Button>
          </div>
        </div>
      </motion.header>

      <main className="flex-grow">
        {children}
      </main>

      <footer className="bg-gray-100 py-8">
        <div className="container mx-auto px-4">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="mb-4 md:mb-0">
              <p className="text-sm text-gray-600">&copy; 2023 ProdCheck. All rights reserved.</p>
            </div>
            <nav className="flex space-x-4">
              <a href="#" className="text-sm text-gray-600 hover:text-blue-600">Terms</a>
              <a href="#" className="text-sm text-gray-600 hover:text-blue-600">Privacy</a>
              <a href="#" className="text-sm text-gray-600 hover:text-blue-600">Contact</a>
            </nav>
          </div>
        </div>
      </footer>
    </div>
  )
}

// Landing Page Component
const LandingPage = () => {
  const { scrollYProgress } = useScroll()
  const opacity = useTransform(scrollYProgress, [0, 0.5], [1, 0])

  return (
    <>
      <section className="relative py-20 overflow-hidden">
        <motion.div 
          className="absolute inset-0 z-0"
          style={{ opacity }}
        >
          <div className="absolute inset-0 bg-blue-500 opacity-10"></div>
          <motion.div
            className="absolute inset-0 bg-gradient-to-r from-blue-400 to-purple-500 opacity-20"
            animate={{
              scale: [1, 1.1, 1],
              rotate: [0, 5, 0],
            }}
            transition={{
              duration: 20,
              repeat: Infinity,
              repeatType: "reverse"
            }}
          ></motion.div>
        </motion.div>
        <div className="container mx-auto px-4 relative z-10">
          <motion.div 
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            className="text-center"
          >
            <h1 className="text-4xl md:text-6xl font-bold text-gray-900 mb-6">
              Protect Your Brand with ProdCheck
            </h1>
            <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
              Monitor and enforce MAP policy violations across e-commerce platforms to maintain brand integrity and fair competition.
            </p>
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ delay: 0.5, duration: 0.8 }}
            >
              <Button size="lg" asChild className="mr-4">
                <Link to="/platform">
                  Get Started
                  <ChevronRight className="ml-2 h-4 w-4" />
                </Link>
              </Button>
              <Button size="lg" variant="outline">
                Watch Demo
              </Button>
            </motion.div>
          </motion.div>
        </div>
      </section>

      <section id="features" className="py-20 bg-white">
        <div className="container mx-auto px-4">
          <motion.h2 
            className="text-3xl font-bold text-center mb-12"
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
          >
            Our Features
          </motion.h2>
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { icon: Globe, title: "Multi-Platform Monitoring", description: "Monitor MAP violations across major e-commerce platforms in real-time." },
              { icon: BarChart, title: "Advanced Analytics", description: "Gain insights with detailed reports and trend analysis." },
              { icon: Shield, title: "Brand Protection", description: "Safeguard your brand's reputation and maintain pricing integrity." },
            ].map((feature, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5, delay: index * 0.1 }}
                viewport={{ once: true }}
                className="bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300"
              >
                <feature.icon className="h-12 w-12 text-blue-600 mb-4" />
                <h3 className="text-xl font-semibold mb-2">{feature.title}</h3>
                <p className="text-gray-600">{feature.description}</p>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      <section id="pricing" className="py-20 bg-gray-50">
        <div className="container mx-auto px-4">
          <motion.h2 
            className="text-3xl font-bold text-center mb-12"
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
          >
            Pricing Plans
          </motion.h2>
          <div className="grid md:grid-cols-3 gap-8">
            {[
              { name: "Basic", price: "$99", features: ["Monitor up to 100 products", "Daily scans", "Email alerts"] },
              { name: "Pro", price: "$299", features: ["Monitor up to 500 products", "Hourly scans", "Email and SMS alerts", "Advanced analytics"] },
              { name: "Enterprise", price: "Custom", features: ["Unlimited product monitoring", "Real-time scans", "Priority support", "Custom integrations"] },
            ].map((plan, index) => (
              <motion.div
                key={index}
                initial={{ opacity: 0, y: 20 }}
                whileInView={{ opacity: 1, y: 0 }}
                transition={{ duration: 0.5, delay: index * 0.1 }}
                viewport={{ once: true }}
                className="bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition-shadow duration-300"
              >
                <h3 className="text-2xl font-bold mb-4">{plan.name}</h3>
                <p className="text-4xl font-bold mb-6">{plan.price}<span className="text-sm font-normal">/month</span></p>
                <ul className="mb-6">
                  {plan.features.map((feature, i) => (
                    <li key={i} className="flex items-center mb-2">
                      <Check className="h-5 w-5 text-green-500 mr-2" />
                      <span>{feature}</span>
                    </li>
                  ))}
                </ul>
                <Button className="w-full" variant={index === 1 ? "default" : "outline"}>Choose Plan</Button>
              </motion.div>
            ))}
          </div>
        </div>
      </section>

      <section id="contact" className="py-20 bg-white">
        <div className="container mx-auto px-4">
          <motion.h2 
            className="text-3xl font-bold text-center mb-12"
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
          >
            Contact Us
          </motion.h2>
          <div className="max-w-lg mx-auto">
            <motion.form 
              className="space-y-4"
              initial={{ opacity: 0, y: 20 }}
              whileInView={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.8 }}
              viewport={{ once: true }}
            >
              <Input placeholder="Your Name" />
              <Input placeholder="Your Email" type="email" />
              <Input placeholder="Subject" />
              <textarea
                className="min-h-[100px] w-full rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50"
                placeholder="Your Message"
              ></textarea>
              <Button type="submit" className="w-full">Send Message</Button>
            </motion.form>
          </div>
        </div>
      </section>

      <section className="py-20 bg-blue-600 text-white">
        <div className="container mx-auto px-4">
          <motion.div 
            className="text-center"
            initial={{ opacity: 0, y: 20 }}
            whileInView={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
            viewport={{ once: true }}
          >
            <h2 className="text-3xl font-bold mb-4">Ready to protect your brand?</h2>
            <p className="text-xl mb-8">Start your free trial today and see the difference ProdCheck can make.</p>
            <Button size="lg" variant="secondary" asChild>
              <Link to="/platform">
                Get Started Now
                <ArrowRight className="ml-2 h-4 w-4" />
              </Link>
            </Button>
          </motion.div>
        </div>
      </section>
    </>
  )
}

// Platform Page Component
const PlatformPage = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false)

  const handleLogin = (e: React.FormEvent) => {
    e.preventDefault()
    // Implement actual login logic here
    setIsLoggedIn(true)
  }

  if (!isLoggedIn) {
    return (
      <div className="min-h-screen bg-gray-50 flex flex-col justify-center py-12 sm:px-6 lg:px-8">
        <div className="sm:mx-auto sm:w-full sm:max-w-md">
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">Sign in to your account</h2>
        </div>

        <div className="mt-8 sm:mx-auto sm:w-full sm:max-w-md">
          <div className="bg-white py-8 px-4 shadow sm:rounded-lg sm:px-10">
            <form className="space-y-6" onSubmit={handleLogin}>
              <div>
                <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                  Email address
                </label>
                <div className="mt-1">
                  <Input id="email" name="email" type="email" autoComplete="email" required />
                </div>
              </div>

              <div>
                <label htmlFor="password" className="block text-sm font-medium text-gray-700">
                  Password
                </label>
                <div className="mt-1">
                  <Input id="password" name="password" type="password" autoComplete="current-password" required />
                </div>
              </div>

              <div>
                <Button type="submit" className="w-full">
                  Sign in
                </Button>
              </div>
            </form>
          </div>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            <h2 className="text-3xl font-bold text-gray-900 mb-6">Welcome to ProdCheck Platform</h2>
            <p className="text-xl text-gray-600 mb-8">
              Our platform provides powerful tools to monitor and enforce MAP policies across various e-commerce channels.
            </p>
          </motion.div>

          <div className="grid md:grid-cols-2 gap-8">
            <motion.div
              initial={{ opacity: 0, x: -20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ duration: 0.8, delay: 0.2 }}
              className="bg-white p-6 rounded-lg shadow-md"
            >
              <h3 className="text-2xl font-semibold mb-4">How It Works</h3>
              <ol className="list-decimal list-inside space-y-2">
                <li>Connect your product catalog</li>
                <li>Set up your MAP policies</li>
                <li>Monitor e-commerce channels in real-time</li>
                <li>Receive alerts for policy violations</li>
                <li>Take action to enforce your policies</li>
              </ol>
            </motion.div>

            <motion.div
              initial={{ opacity: 0, x: 20 }}
              animate={{ opacity: 1, x: 0 }}
              transition={{ duration: 0.8, delay: 0.4 }}
              className="bg-white p-6 rounded-lg shadow-md"
            >
              <h3 className="text-2xl font-semibold mb-4">Key Features</h3>
              <ul className="space-y-2">
                <li className="flex items-center">
                  <Globe className="h-5 w-5 text-blue-600 mr-2" />
                  Multi-platform monitoring
                </li>
                <li className="flex items-center">
                  <BarChart className="h-5 w-5 text-blue-600 mr-2" />
                  Advanced analytics and reporting
                </li>
                <li className="flex items-center">
                  <AlertTriangle className="h-5 w-5 text-blue-600 mr-2" />
                  Real-time violation alerts
                </li>
                <li className="flex items-center">
                  <Shield className="h-5 w-5 text-blue-600 mr-2" />
                  Automated enforcement actions
                </li>
              </ul>
            </motion.div>
          </div>

          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.6 }}
            className="mt-12 text-center"
          >
            <p className="text-xl text-gray-600 mb-4">Ready to start protecting your brand?</p>
            <Button size="lg" asChild>
              <Link to="/analyze">
                Analyze Market
                <ChevronRight className="ml-2 h-4 w-4" />
              </Link>
            </Button>
          </motion.div>
        </div>
      </div>
    </div>
  )
}

// Analyze Market Page Component
const AnalyzePage = () => {
  const [searchTerm, setSearchTerm] = useState('')
  const [selectedPlatform, setSelectedPlatform] = useState('')
  const [priceRange, setPriceRange] = useState('')
  const [searchResults, setSearchResults] = useState(null)

  const handleSearch = async (e: React.FormEvent) => {
    e.preventDefault()
    // Implement actual search logic here, connecting to your backend
    // For now, we'll simulate a search result
    setSearchResults({
      website: searchTerm,
      found: Math.random() > 0.5,
      pdfUrl: '/sample-report.pdf'
    })
  }

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            <h2 className="text-3xl font-bold text-gray-900 mb-6">Analyze E-commerce Platforms</h2>
            <p className="text-xl text-gray-600 mb-8">
              Search for specific e-commerce websites or use filters to find platforms that match your criteria.
            </p>
          </motion.div>

          <div className="bg-white p-6 rounded-lg shadow-md">
            <form onSubmit={handleSearch} className="space-y-4">
              <div className="flex space-x-4">
                <div className="flex-grow">
                  <label htmlFor="search" className="block text-sm font-medium text-gray-700 mb-1">
                    Search E-commerce Website
                  </label>
                  <div className="relative">
                    <Input
                      id="search"
                      type="text"
                      placeholder="Enter website name"
                      value={searchTerm}
                      onChange={(e) => setSearchTerm(e.target.value)}
                      className="pl-10"
                    />
                    <Search className="absolute left-3 top-1/2 transform -translate-y-1/2 h-5 w-5 text-gray-400" />
                  </div>
                </div>
                <div>
                  <label htmlFor="platform" className="block text-sm font-medium text-gray-700 mb-1">
                    Platform
                  </label>
                  <Select
                    id="platform"
                    value={selectedPlatform}
                    onChange={(e) => setSelectedPlatform(e.target.value)}
                  >
                    <option value="">All Platforms</option>
                    <option value="amazon">Amazon</option>
                    <option value="ebay">eBay</option>
                    <option value="walmart">Walmart</option>
                  </Select>
                </div>
                <div>
                  <label htmlFor="price" className="block text-sm font-medium text-gray-700 mb-1">
                    Price Range
                  </label>
                  <Select
                    id="price"
                    value={priceRange}
                    onChange={(e) => setPriceRange(e.target.value)}
                  >
                    <option value="">Any Price</option>
                    <option value="0-50">$0 - $50</option>
                    <option value="51-100">$51 - $100</option>
                    <option value="101+">$101+</option>
                  </Select>
                </div>
              </div>
              <Button type="submit" className="w-full">
                <Filter className="mr-2 h-4 w-4" />
                Search and Filter
              </Button>
            </form>
          </div>

          {searchResults && (
            <motion.div
              initial={{ opacity: 0, y: 20 }}
              animate={{ opacity: 1, y: 0 }}
              transition={{ duration: 0.8 }}
              className="mt-8 bg-white p-6 rounded-lg shadow-md"
            >
              <h3 className="text-2xl font-semibold mb-4">Search Results</h3>
              {searchResults.found ? (
                <div>
                  <p className="text-green-600 mb-4">Website found in our database!</p>
                  <Button asChild>
                    <a href={searchResults.pdfUrl} target="_blank" rel="noopener noreferrer">
                      <BarChart className="mr-2 h-4 w-4" />
                      View Market Analysis Report (PDF)
                    </a>
                  </Button>
                </div>
              ) : (
                <p className="text-red-600">Website not found in our database. Try another search or use the filters.</p>
              )}
            </motion.div>
          )}
        </div>
      </div>
    </div>
  )
}

// Dashboard Page Component
const DashboardPage = () => {
  const [notifications, setNotifications] = useState([])

  useEffect(() => {
    // Simulating notifications from backend
    const fetchNotifications = async () => {
      // In a real application, this would be an API call
      const mockNotifications = [
        { id: 1, type: 'violation', message: 'MAP violation detected on Amazon for Product A' },
        { id: 2, type: 'price_change', message: 'Price change detected for Product B on eBay' },
        { id: 3, type: 'new_listing', message: 'New listing found for Product C on Walmart' },
      ]
      setNotifications(mockNotifications)
    }

    fetchNotifications()
  }, [])

  return (
    <div className="min-h-screen bg-gray-50">
      <div className="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div className="px-4 py-6 sm:px-0">
          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8 }}
          >
            <h2 className="text-3xl font-bold text-gray-900 mb-6">Your ProdCheck Dashboard</h2>
          </motion.div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">Total Monitored Products</CardTitle>
                <BarChart className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">1,234</div>
                <p className="text-xs text-muted-foreground">+20% from last month</p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">Active MAP Violations</CardTitle>
                <AlertTriangle className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">23</div>
                <p className="text-xs text-muted-foreground">-5% from last week</p>
              </CardContent>
            </Card>
            <Card>
              <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
                <CardTitle className="text-sm font-medium">Estimated Revenue Impact</CardTitle>
                <DollarSign className="h-4 w-4 text-muted-foreground" />
              </CardHeader>
              <CardContent>
                <div className="text-2xl font-bold">$12,345</div>
                <p className="text-xs text-muted-foreground">Based on current violations</p>
              </CardContent>
            </Card>
          </div>

          <motion.div
            initial={{ opacity: 0, y: 20 }}
            animate={{ opacity: 1, y: 0 }}
            transition={{ duration: 0.8, delay: 0.2 }}
          >
            <h3 className="text-2xl font-semibold mb-4">Recent Notifications</h3>
            <Card>
              <CardContent className="p-0">
                <ul className="divide-y divide-gray-200">
                  {notifications.map((notification) => (
                    <li key={notification.id} className="p-4 hover:bg-gray-50">
                      <div className="flex items-center space-x-4">
                        <div className="flex-shrink-0">
                          {notification.type === 'violation' && (
                            <AlertTriangle className="h-6 w-6 text-red-500" />
                          )}
                          {notification.type === 'price_change' && (
                            <DollarSign className="h-6 w-6 text-yellow-500" />
                          )}
                          {notification.type === 'new_listing' && (
                            <Bell className="h-6 w-6 text-green-500" />
                          )}
                        </div>
                        <div className="flex-1 min-w-0">
                          <p className="text-sm font-medium text-gray-900 truncate">
                            {notification.message}
                          </p>
                        </div>
                        <div>
                          <Button variant="outline" size="sm">View</Button>
                        </div>
                      </div>
                    </li>
                  ))}
                </ul>
              </CardContent>
            </Card>
          </motion.div>
        </div>
      </div>
    </div>
  )
}

// Main App Component
export function ProdCheckAppComponent() {
  return (
    <Router>
      <Layout>
        <Routes>
          <Route path="/" element={<LandingPage />} />
          <Route path="/platform" element={<PlatformPage />} />
          <Route path="/analyze" element={<AnalyzePage />} />
          <Route path="/dashboard" element={<DashboardPage />} />
        </Routes>
      </Layout>
    </Router>
  )
}