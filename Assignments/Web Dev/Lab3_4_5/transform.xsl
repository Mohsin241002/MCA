<?xml version="1.0" encoding="UTF-8"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform"
                xmlns:xs="http://www.w3.org/2001/XMLSchema"
                xmlns:fn="http://www.w3.org/2005/xpath-functions"
                xmlns:my="http://www.example.com/games">

    <!-- Define the output format -->
    <xsl:output method="html" indent="yes"/>
    
    <!-- Template for matching the root element -->
    <xsl:template match="my:games">
        <html>
            <head>
                <meta charset="UTF-8"/>
                <title>Gaming Data</title>
                <link rel="stylesheet" href="css/style.css"/>
                <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/css/bootstrap.min.css" rel="stylesheet"/>
            </head>
            <body>
                
                <main class="container mt-4">
                    <table class="table table-bordered">
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Name</th>
                                <th>Type</th>
                                <th>Release Date</th>
                                <th>Rating</th>
                                <th>Platform</th>
                                <th>Developer</th>
                                <th>Price (Currency)</th>
                            </tr>
                        </thead>
                        <tbody>
                            <!-- Iterate over each game element -->
                            <xsl:for-each select="my:game">
                            <tr>
                            <td><xsl:value-of select="my:id"/></td>
                                    <td><xsl:value-of select="my:name"/></td>
                                    <td><xsl:value-of select="my:type"/></td>
                                    <td><xsl:value-of select="my:releaseDate"/></td>
                                    <td><xsl:value-of select="my:rating"/></td>
                                    <td><xsl:value-of select="my:platform"/></td>
                                    <td><xsl:value-of select="my:developer"/></td>
                                    <td><xsl:value-of select="concat(my:price, ' (', @currency, ')')"/></td>
                                </tr>
                            </xsl:for-each>
                        </tbody>
                    </table>
                </main>
                
                <script src="https://code.jquery.com/jquery-3.5.1.slim.min.js"></script>
                <script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.9.2/dist/umd/popper.min.js"></script>
                <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.5.2/js/bootstrap.min.js"></script>
                <!-- Font Awesome JS -->
                <script src="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/js/all.min.js"></script>
            </body>
        </html>
    </xsl:template>
</xsl:stylesheet>