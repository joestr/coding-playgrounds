package org.example.kotlinplayground

import org.eclipse.jetty.server.Connector
import org.eclipse.jetty.server.Server
import org.eclipse.jetty.server.ServerConnector


fun main(args: Array<String>) {
    var server = Server()
    val connector = ServerConnector(server)
    connector.port = 8090
    server.setConnectors(arrayOf<Connector>(connector))
}

fun main2() {
    val name = "Kotlin"
    println("Hello, " + name + "!")

    for (i in 1..5) {
        println("i = $i")
    }
}