// Native Angular Services
import { Component, OnInit, Renderer2 } from '@angular/core';
import { Title } from '@angular/platform-browser';
import * as animate from 'animejs';
// declare let animejs: any;

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [ Title ],
})
export class AppComponent implements OnInit {
  title = 'kvmac.io';
  fontColor: any;
  bgColor1: any;
  bgColor2: any;
  bgColorButton1: any;
  bgColorButton2: any;
  // anime = new animejs;
  // cssProperties: any;
  // anime: any;

  constructor(
    private titleService: Title,
    private renderer: Renderer2,
    ) {
      // this.anime = new animejs;
      // this.titleService.setTitle(this.title);
      this.setTitle();

    // this.cssProperties = this.anime.anime({
    //     targets: '#cssProperties .el',
    //     left: '240px',
    //     backgroundColor: '#FFF',
    //     borderRadius: 40,
    //     easing: 'easeInOutQuad'
    //   });


      this.fontColor = this.getRandomColor();
      this.bgColor1 = this.getRandomColor();
      this.bgColor2 = this.getRandomColor();
      this.bgColorButton1 = this.getRandomColor();
      this.bgColorButton2 = this.getRandomColor();
    }

          // let body = document.getElementsByTagName('body')[0];
    // body.classList.add("background");
  // public setTitle( newTitle: string) {
  //   this.titleService.setTitle(newTitle)
  // }
  private setTitle() {
    this.titleService.setTitle(this.title);
  }
  public animateSpin() {
    animate({
      targets: '#specificUnitValue .el',
      translateX: '20em',
      rotate: '1.5turn'
    });
  }
  public animateSlide() {
    animate({
      targets: '#alternate .el',
      translateX: 250,
      direction: 'alternate'
    });
  }
  // setTitle(title);
  getRandomColor() {
    const letters = '0123456789ABCDEF';
    let color = '#';
    for (let i = 0; i < 6; i++ ) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
  }

  ngOnInit(): void {
  }

}
