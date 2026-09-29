"""Shared SVG primitives for licence teaching diagrams."""
import textwrap
import math
from html import escape
from pathlib import Path

ROOT = Path(__file__).resolve().parents[2]
ASSETS = ROOT / 'web/frontend-react/public/illustrations'
CYAN, AMBER, PURPLE = '#58dfed', '#ffcd70', '#bbabff'
WHITE, MUTED, RED, PANEL = '#edf5ff', '#a9bdd3', '#ff8e9c', '#14253b'


class SVG:
    def __init__(self, module, title, subtitle, description, program="PPL"):
        self.parts = [f'<svg xmlns="http://www.w3.org/2000/svg" width="1200" height="720" viewBox="0 0 1200 720" role="img" aria-labelledby="title desc"><title id="title">{escape(title)}</title><desc id="desc">{escape(description)}</desc>',
                      '<defs><linearGradient id="bg" x2="1" y2="1"><stop stop-color="#10243d"/><stop offset="1" stop-color="#080f1e"/></linearGradient></defs>']
        self.rect(0, 0, 1200, 720, 'url(#bg)', radius=24)
        self.text(42, 45, f'AEROPATH / {program} / COMPRENDRE EN IMAGE', 14, CYAN, weight=700)
        self.text(42, 96, title, 34, WHITE, weight=700)
        self.text(42, 135, subtitle, 19, MUTED)
        self.rect(32, 170, 1136, 450, PANEL, '#2b415d', 20)
        self.text(42, 668, module, 14, CYAN, weight=700)
        self.text(1158, 668, 'Schéma pédagogique · Données d’exercice, sauf mention contraire', 14, MUTED, anchor='end')
        self.text(42, 697, 'Relier le schéma aux explications et aux références de la leçon.', 14, MUTED)

    def rect(self, x, y, w, h, fill, stroke='none', radius=0):
        self.parts.append(f'<rect x="{x}" y="{y}" width="{w}" height="{h}" rx="{radius}" fill="{fill}" stroke="{stroke}"/>')

    def text(self, x, y, text, size=20, color=WHITE, anchor='start', weight=400):
        self.parts.append(f'<text x="{x}" y="{y}" fill="{color}" font-family="Segoe UI, Arial, sans-serif" font-size="{size}" font-weight="{weight}" text-anchor="{anchor}">{escape(str(text))}</text>')

    def wrap(self, x, y, text, width=270, size=20, color=MUTED, line=29, weight=400):
        lines = textwrap.wrap(text, max(8, int(width / (size * .55))), break_long_words=False)
        for i, value in enumerate(lines):
            self.text(x, y + i * line, value, size, color, weight=weight)
        return y + len(lines) * line

    def line(self, x1, y1, x2, y2, color=MUTED, width=2, dashed=False):
        self.parts.append(f'<path d="M{x1} {y1}L{x2} {y2}" fill="none" stroke="{color}" stroke-width="{width}"'+(' stroke-dasharray="7 7"' if dashed else '')+'/>')

    def path(self, d, fill='none', stroke=CYAN, width=3, opacity=1):
        self.parts.append(f'<path d="{d}" fill="{fill}" stroke="{stroke}" stroke-width="{width}" stroke-linejoin="round" opacity="{opacity}"/>')

    def circle(self, x, y, r, fill=PANEL, stroke=CYAN, width=2):
        self.parts.append(f'<circle cx="{x}" cy="{y}" r="{r}" fill="{fill}" stroke="{stroke}" stroke-width="{width}"/>')

    def arrow(self, x1, y1, x2, y2, color=CYAN, width=4):
        self.line(x1, y1, x2, y2, color, width)
        a = math.atan2(y2-y1, x2-x1)
        points = [(x2, y2), (x2-15*math.cos(a-.45), y2-15*math.sin(a-.45)), (x2-15*math.cos(a+.45), y2-15*math.sin(a+.45))]
        self.path('M'+'L'.join(f'{x:.1f} {y:.1f}' for x,y in points)+'Z', color, color, 1)

    def plane(self, x, y, scale=1, rotation=0, color=WHITE):
        self.parts.append(f'<g transform="translate({x} {y}) rotate({rotation}) scale({scale})">')
        self.path('M0-80C-12-74-13-45-10-20L-85 17-85 30-10 15-6 58-30 75-30 83 0 74 30 83 30 75 6 58 10 15 85 30 85 17 10-20C13-45 12-74 0-80Z', color, '#7895b3', 2)
        self.path('M-6-42Q0-51 6-42L6-25H-6Z', '#18324c', '#18324c', 1)
        self.parts.append('</g>')

    def cloud(self, x, y, scale=1, color='#9fb7d0'):
        self.parts.append(f'<g transform="translate({x} {y}) scale({scale})">')
        self.path('M-60 15C-88 10-78-25-52-25C-52-58-9-64 5-39C35-62 65-38 60-13C91-12 93 19 67 23H-60Z', color, color, 1)
        self.parts.append('</g>')

    def icon(self, name, x, y, color=CYAN):
        if name == 'plane':
            self.plane(x, y, .48, 30, color)
        elif name == 'cloud':
            self.cloud(x, y+8, .55, color)
        elif name == 'clock':
            self.circle(x, y, 35, PANEL, color, 3)
            self.line(x,y,x,y-23,color,4); self.line(x,y,x+20,y+12,color,4)
        elif name == 'globe':
            self.circle(x,y,37,PANEL,color,3)
            self.parts.append(f'<ellipse cx="{x}" cy="{y}" rx="16" ry="37" fill="none" stroke="{color}" stroke-width="2"/>')
            self.line(x-35,y,x+35,y,color,2)
            self.line(x-29,y-20,x+29,y-20,color,2); self.line(x-29,y+20,x+29,y+20,color,2)
        elif name == 'radio':
            self.rect(x-25,y-20,50,55,PANEL,color,7)
            self.line(x+12,y-20,x+25,y-51,color,3)
            self.rect(x-14,y-10,28,12,color,radius=2)
            for dy in [13,22]: self.line(x-13,y+dy,x+13,y+dy,color,2)
            self.path(f'M{x+37} {y-25}q25 20 0 40',stroke=color)
        elif name == 'eye':
            self.path(f'M{x-45} {y}Q{x} {y-45} {x+45} {y}Q{x} {y+45} {x-45} {y}Z',stroke=color)
            self.circle(x,y,14,color,color)
        elif name == 'fuel':
            self.rect(x-30,y-33,48,68,PANEL,color,5)
            self.rect(x-22,y-23,32,20,color,radius=3)
            self.path(f'M{x+18} {y-10}H{x+32}V{y+22}Q{x+48} {y+35} {x+48} {y+16}V{y-12}L{x+35} {y-26}',stroke=color)
        elif name == 'map':
            self.path(f'M{x-40} {y-30}L{x-12} {y-40}L{x+12} {y-27}L{x+40} {y-37}V{y+27}L{x+12} {y+38}L{x-12} {y+25}L{x-40} {y+35}Z',stroke=color)
            self.line(x-12,y-40,x-12,y+25,color); self.line(x+12,y-27,x+12,y+38,color)
        elif name == 'shield':
            self.path(f'M{x} {y-40}L{x+35} {y-24}V{y+4}Q{x+30} {y+28} {x} {y+43}Q{x-30} {y+28} {x-35} {y+4}V{y-24}Z',stroke=color)
            self.path(f'M{x-17} {y}L{x-3} {y+15}L{x+21} {y-15}',stroke=color,width=4)
        elif name == 'wind':
            for dy,dx in [(-20,0),(0,15),(20,-8)]:
                self.path(f'M{x-42} {y+dy}H{x+20+dx}q22 0 16-16',stroke=color)
        else:
            self.rect(x-30,y-38,60,76,PANEL,color,6)
            for dy in [-17,0,17]: self.line(x-17,y+dy,x+17,y+dy,color,3)

    def rail(self, items):
        self.line(829,205,829,583,'#304961',1)
        for i,(heading,body) in enumerate(items):
            y=230+i*118; color=[CYAN,AMBER,PURPLE][i%3]
            self.circle(864,y-6,5,color,color)
            self.text(883,y,heading,20,color,weight=700)
            self.wrap(855,y+32,body,270,18,line=26)

    def finish(self):
        return '\n'.join(self.parts+['</svg>'])
